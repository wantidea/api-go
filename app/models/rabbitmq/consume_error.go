package rabbitmq

type ConsumeError struct {
	Id             int64  `json:"id" gorm:"primaryKey"`
	ExchangeName   string `json:"exchange_name"`
	ExchangeType   string `json:"exchange_type"`
	QueueName      string `json:"queue_name"`
	RoutingKey     string `json:"routing_key"`
	MessageContent string `json:"message_content"`
	ErrorContent   string `json:"error_content"`
	Ack            int64  `json:"ack"`
	CreatedAt      int64  `json:"created_at"`
	UpdatedAt      int64  `json:"updated_at"`
}

func (m *ConsumeError) TableName() string {
	return "rabbitmq_consume_error"
}
