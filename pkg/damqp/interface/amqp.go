package brokerInterface

type RabbitMQ interface {
	StartingConsumeMQ(queueName string, consumeFunc func(d []byte) error) error
	PublishMQ(queueName, data, urlMQ string) error
}
