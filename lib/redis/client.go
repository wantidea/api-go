package redis

import (
	"api-go/lib/config"
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var rdb *redis.Client

func Setup() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisConfig.Host + ":" + config.RedisConfig.Port,
		Password: config.RedisConfig.Password,
		DB:       config.RedisConfig.DB,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Redis 连接失败: %v", err)
	}
}

func Rdb() *redis.Client {
	return rdb
}
