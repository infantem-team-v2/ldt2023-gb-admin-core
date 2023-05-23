package brokerInterface

type Kafka interface {
	Publish(brokers []string, topic string, message string) error
	Consume(brokers []string, topic string, groupID string) error
}
