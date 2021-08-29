package main

import (
	SystemServices "api-go/app/services/system"
	"api-go/lib/config"
	"api-go/lib/logger"
	"api-go/lib/mongodb"
	"api-go/lib/orm"
	"api-go/lib/redis"
	"api-go/lib/request"
	"api-go/routers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	config.Setup()

	if config.AppConfig.MysqlState {
		orm.Setup()
	}

	if config.AppConfig.RedisState {
		redis.Setup()
	}

	if config.AppConfig.MongodbState {
		mongodb.Setup()
	}

	if config.AppConfig.LoggerState {
		logger.Setup()
	}

	if config.AppConfig.RabbitmqState {
		SystemServices.ListenRefreshTreeMenuPromise()
	}

	if err := request.InitTrans(config.AppConfig.Locale); err != nil {
		log.Fatalf("request 中文初始化失败： %s\n", err)
	}
}

func main() {
	// 退出关闭 mysql 连接
	if config.AppConfig.MysqlState {
		defer orm.CloseDB()
	}

	port := fmt.Sprintf(":%s", config.AppConfig.Port)
	router := routers.InitRouter()
	readTimeout := config.AppConfig.ReadTimeout
	writeTimeout := config.AppConfig.WriteTimeout
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		// 服务连接
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("监听: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("关闭服务器...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("服务器关闭:", err)
	}
	log.Println("服务器已关闭")
}
