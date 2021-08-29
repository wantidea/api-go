package mongodb

import (
	"api-go/lib/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var client *mongo.Client

func Setup() {
	mongodbUrl := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		config.MongodbConfig.Connection,
		config.MongodbConfig.UserName,
		config.MongodbConfig.Password,
		config.MongodbConfig.Host,
		config.MongodbConfig.Port,
		config.MongodbConfig.Database)

	clientOptions := options.Client().ApplyURI(mongodbUrl)

	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("mongodb 数据库连接失败: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("mongodb 数据库连接失败: %v", err)
	}
}

func CloseMongoDB() {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("mongodb 数据库关闭失败: %v", err)
	}
}
