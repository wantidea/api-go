package rabbitmq

import "github.com/streadway/amqp"

type PublishContent interface {
	MsgContent() string
}

type Producer struct {
	QueueName    string
	ExchangeName string
	ExchangeType string
	RoutingKey   string
}

func (p *Producer) Publish(publishItem PublishContent) error {
	var err error
	r := &RabbitMQ{
		queueName:    p.QueueName,
		exchangeName: p.ExchangeName,
		exchangeType: p.ExchangeType,
		routingKey:   p.RoutingKey,
	}

	if r.channel == nil {
		err = r.Connect()
		if err != nil {
			return err
		}
	}
	defer r.Close()

	err = r.channel.ExchangeDeclare(
		r.exchangeName,
		r.exchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		failOnError(err, "生产者 交换机获取失败！")
		return err
	}

	err = r.channel.Publish(
		r.exchangeName,
		r.routingKey,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(publishItem.MsgContent()),
		})
	if err != nil {
		failOnError(err, "消息发布失败！")
		return err
	}

	return nil
}
