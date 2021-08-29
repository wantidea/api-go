package rabbitmq

import (
	"api-go/lib/config"
	"api-go/lib/logger"
	"fmt"
	"github.com/streadway/amqp"
	"sync"
)

// 连接 RabbitMQ
type RabbitMQ struct {
	connection   *amqp.Connection
	channel      *amqp.Channel
	queueName    string // 队列名称
	exchangeName string // 交换机名称
	exchangeType string // 交换机类型
	routingKey   string // 路由名称
	rabbitUrl    string // 连接信息
	mu           sync.RWMutex
}

// 定义队列交换机对象
type QueueExchange struct {
	QueueName    string // 队列名称
	ExchangeName string // 交换机名称
	RoutingKey   string // 路由名称
}

//创建一个新的操作对象
func New(q *QueueExchange) *RabbitMQ {
	return &RabbitMQ{
		queueName:    q.QueueName,
		exchangeName: q.ExchangeName,
		routingKey:   q.RoutingKey,
	}
}

//创建连接
func (r *RabbitMQ) Connect() error {
	var err error
	r.rabbitUrl = fmt.Sprintf("amqp://%s:%s@%s:%s/",
		config.RabbitMQConfig.UserName,
		config.RabbitMQConfig.Password,
		config.RabbitMQConfig.Host,
		config.RabbitMQConfig.Port)

	r.connection, err = amqp.Dial(r.rabbitUrl)
	if err != nil {
		failOnError(err, "创建连接错误！")
		return err
	}

	r.channel, err = r.connection.Channel()
	if err != nil {
		failOnError(err, "打开 channel 失败！")
		return err
	}

	return nil
}

// Close 关闭连接
func (r *RabbitMQ) Close() {
	_ = r.channel.Close()
	_ = r.connection.Close()
}

// failOnError 查看每个 amqp 函数调用返回值
func failOnError(err error, msg string) {
	if err != nil {
		logger.Debug(msg, err)
	}
}
