package rabbitmq

import (
	RabbitMQModels "api-go/app/models/rabbitmq"
	"api-go/lib/orm"
	"fmt"
)

type ConsumeContent interface {
	MsgConsume([]byte) error
}

type Receiver struct {
	QueueName    string
	ExchangeName string
	ExchangeType string
	RoutingKey   string
	AckCount     int
}

// Consume 监听消费队列
func (r *Receiver) Consume(consumeItem ConsumeContent) {
	var err error
	rabbitMQ := &RabbitMQ{
		queueName:    r.QueueName,
		exchangeName: r.ExchangeName,
		exchangeType: r.ExchangeType,
		routingKey:   r.RoutingKey,
	}

	if rabbitMQ.channel == nil {
		err = rabbitMQ.Connect()
		if err != nil {
			return
		}
	}
	defer rabbitMQ.Close()

	err = rabbitMQ.channel.Qos(
		10, // 预计数量
		0,  // 预计大小
		false,
	)
	if err != nil {
		failOnError(err, "队列控制失败！")
		return
	}

	err = rabbitMQ.channel.ExchangeDeclare(
		rabbitMQ.exchangeName,
		rabbitMQ.exchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		failOnError(err, "消费者 交换机获取失败！")
		return
	}

	_, err = rabbitMQ.channel.QueueDeclare(
		rabbitMQ.queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		failOnError(err, "消费者队列获取失败！")
		return
	}

	err = rabbitMQ.channel.QueueBind(
		rabbitMQ.queueName,
		rabbitMQ.routingKey,
		rabbitMQ.exchangeName,
		false,
		nil,
	)
	if err != nil {
		failOnError(err, "消费者队列绑定失败")
		return
	}

	messageList, err := rabbitMQ.channel.Consume(
		rabbitMQ.queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		failOnError(err, "消费失败！")
		return
	}

	forever := make(chan bool)
	go func() {
		for message := range messageList {
			if r.AckCount == 0 {
				r.AckCount = 3
			}

			var ackFlag bool // 消费成功 Flag
			for i := 0; i < r.AckCount; i++ {
				err = consumeItem.MsgConsume(message.Body)
				if err != nil {
					_ = message.Ack(true)
					ackFlag = false
				} else {
					err = message.Ack(false)
					if err != nil {
						ackFlag = false
					} else {
						ackFlag = true
					}
					break
				}
			}

			// 消费失败记录
			if !ackFlag {
				model := &RabbitMQModels.ConsumeError{
					ExchangeName:   r.ExchangeName,
					ExchangeType:   r.ExchangeType,
					QueueName:      r.QueueName,
					RoutingKey:     r.RoutingKey,
					MessageContent: string(message.Body),
					ErrorContent:   fmt.Sprintln(err),
					Ack:            0,
				}
				res := orm.DB().Create(model)
				if res.Error != nil {
					failOnError(err, "消费失败，记录失败！")
				}
			}
		}
	}()
	<-forever

}
