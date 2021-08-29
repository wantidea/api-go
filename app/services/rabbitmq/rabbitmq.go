package rabbitmq

import (
	"api-go/lib/rabbitmq"
	"github.com/streadway/amqp"
)

// 队列信息
const (
	// 交换机类型
	ExchangeDirect  = amqp.ExchangeDirect
	ExchangeFanout  = amqp.ExchangeFanout
	ExchangeTopic   = amqp.ExchangeTopic
	ExchangeHeaders = amqp.ExchangeDirect

	// 交换机名称
	ERedis = "ERedis" // redis

	// 路由键
	RRedisTreeMenu = "refreshTreeMenu" // 更新树形菜单

	// 消息
	MRedisTreeMenuRefresh = "refresh" // 刷新
)

type Publish struct {
	message string
}

func (t *Publish) MsgContent() string {
	return t.message
}

// RegisterProducer 注册生产者
func RegisterProducer(exchange, exchangeType, routingKey, message string) error {
	producer := &rabbitmq.Producer{
		QueueName:    routingKey,
		ExchangeName: exchange,
		ExchangeType: exchangeType,
		RoutingKey:   routingKey,
	}
	p := &Publish{
		message,
	}
	return producer.Publish(p)
}

// RegisterReceiver 消费
func RegisterReceiver(exchange, exchangeType, routingKey string, c rabbitmq.ConsumeContent) {
	receiver := &rabbitmq.Receiver{
		QueueName:    routingKey,
		ExchangeName: exchange,
		ExchangeType: exchangeType,
		RoutingKey:   routingKey,
	}
	go receiver.Consume(c)
}
