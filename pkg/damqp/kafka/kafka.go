package kafka

import (
	"context"
	"fmt"
	"gb-auth-gate/config"
	"gb-auth-gate/pkg/tlogger"
	"github.com/Shopify/sarama"
	"github.com/sarulabs/di"
)

type Kafka struct {
	cfg    config.Config
	logger tlogger.ILogger
}

func BuildKafka(ctn di.Container) (interface{}, error) {
	return &Kafka{
		cfg:    ctn.Get("config").(config.Config),
		logger: ctn.Get("logger").(tlogger.ILogger),
	}, nil
}

type MessageHandler struct{}

// Реализация интерфейса sarama.ConsumerGroupHandler для обработки сообщений Kafka
func (h *MessageHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h *MessageHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (k Kafka) Consume(brokers []string, topic string, groupID string) error {
	// Создаем конфигурацию для потребителя Kafka
	config := sarama.NewConfig()

	// Устанавливаем начальное смещение на более старое сообщение
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// Создаем клиент Kafka
	client, err := sarama.NewClient(brokers, config)
	if err != nil {
		return fmt.Errorf("failed to create Kafka client: %w", err)
	}

	// Создаем группу потребителей Kafka
	consumerGroup, err := sarama.NewConsumerGroupFromClient(groupID, client)
	if err != nil {
		return fmt.Errorf("failed to create Kafka consumer group: %w", err)
	}

	// Запускаем обработчик сообщений Kafka
	messageHandler := &MessageHandler{}
	ctx := context.Background()
	for {
		if err := consumerGroup.Consume(ctx, []string{topic}, messageHandler); err != nil {
			return fmt.Errorf("failed to consume from Kafka: %w", err)
		}
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
}

func (h *MessageHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		// Обработка сообщения
		fmt.Printf("Consumed message: topic=%s partition=%d offset=%d value=%s\n",
			message.Topic, message.Partition, message.Offset, string(message.Value))

		// Отмечаем сообщение как обработанное
		sess.MarkMessage(message, "")
	}
	return nil
}

func (k Kafka) Publish(brokers []string, topic string, message string) error {
	// Создаем конфигурацию для производителя Kafka
	config := sarama.NewConfig()
	// Устанавливаем подтверждение доставки сообщений
	config.Producer.RequiredAcks = sarama.WaitForAll

	// Создаем клиент Kafka
	client, err := sarama.NewClient(brokers, config)
	if err != nil {
		return fmt.Errorf("failed to create Kafka client: %w", err)
	}
	defer func() {
		if err = client.Close(); err != nil {
			k.logger.Errorf("failed to close Kafka client: ", err)
		}
	}()

	// Создаем производителя Kafka
	producer, err := sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		return fmt.Errorf("failed to create Kafka producer: %w", err)
	}
	defer func() {
		if err = producer.Close(); err != nil {
			k.logger.Errorf("failed to close Kafka producer:", err)
		}
	}()

	// Создаем сообщение Kafka
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	// Отправляем сообщение в Kafka
	producer.Input() <- msg

	// Обрабатываем ошибки при отправке сообщения
	select {
	case <-producer.Successes():
		return nil
	case err = <-producer.Errors():
		return fmt.Errorf("failed to publish message to Kafka: %w", err)
	}
}
