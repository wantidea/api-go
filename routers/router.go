package routers

import (
	"api-go/app/middlewares"
	"api-go/lib/config"
	"api-go/lib/file"
	"api-go/lib/logger"
	"api-go/routers/business/top"
	"api-go/routers/client"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func InitRouter() *gin.Engine {
	gin.SetMode(config.AppConfig.RunMode)

	gin.DisableConsoleColor()

	gin.DefaultWriter = io.MultiWriter(logger.F, os.Stdout)
	//r := gin.Default()
	r := gin.New()
	//r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 中间件
	//r.Use(middlewares.Recover())   // 捕获全局错误 上线开启
	r.Use(middlewares.Cors()) // 跨域设置

	// 加载模块路由
	// B端
	top.AddTopRoute(r) // 总站

	// C端
	client.AddBlogRoute(r) // 博客

	// 静态文件访问
	r.StaticFS(config.ImageConfig.StaticUrl, http.Dir(file.ImageStaticDir()))

	return r
}
