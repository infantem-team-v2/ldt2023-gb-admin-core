package rabbit

import (
	"context"
	"fmt"
	"gb-auth-gate/config"
	"gb-auth-gate/pkg/tlogger"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sarulabs/di"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ConsumerMQ struct {
	cfg    config.Config
	logger tlogger.ILogger
}

func BuildRabbitMQ(ctn di.Container) (interface{}, error) {
	return &ConsumerMQ{
		cfg:    ctn.Get("config").(config.Config),
		logger: ctn.Get("logger").(tlogger.ILogger),
	}, nil
}

func (c ConsumerMQ) StartingConsumeMQ(queueName string, consumeFunc func(d []byte) error) error {
	var queueMap = map[string]string{
		"example": fmt.Sprintf("amqp://%s:%s@rabbitmq-%s:%s/",
			c.cfg.AmqpConfig.RabbitMQ.UserName,
			c.cfg.AmqpConfig.RabbitMQ.Password,
			c.cfg.AmqpConfig.RabbitMQ.Host,
			c.cfg.AmqpConfig.RabbitMQ.Port,
		),
	}
	err := c.consume(queueName, queueMap[queueName], consumeFunc)
	if err != nil {
		return err
	}
	return nil
}

func (c ConsumerMQ) consume(queueName, urlMQ string,
	consumeFunc func(d []byte) error) error {
	conn, err := amqp.Dial(urlMQ)
	if err != nil {
		c.logger.Errorf("Error with connection to RabbitMQ: ", err)
		return err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			c.logger.Errorf("Error with closing consume RabbitMQ: ", err)
		}
	}()
	ch, err := conn.Channel()
	if err != nil {
		c.logger.Errorf("Error with creating channel for consume RabbitMQ: ", err)
		return err
	}
	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil)
	c.logger.Errorf("Error with register consumer: %s", err)

	go func() {
		for msg := range msgs {
			err = consumeFunc(msg.Body)
			if err != nil {
				c.logger.Errorf("", err)
			}
		}
	}()

	// block main thread - wait for shutdown signal
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("stopping consumer")
	return nil
}

func (c ConsumerMQ) PublishMQ(queueName, data, urlMQ string) error {
	conn, err := amqp.Dial(urlMQ)
	defer func() {
		err = conn.Close()
		if err != nil {
			c.logger.Errorf("Error with closing publish RabbitMQ: ", err)
		}
	}()
	if err != nil {
		c.logger.Errorf("Error with connection to RabbitMQ: ", err)
		return err
	}

	ch, err := conn.Channel()
	err = ch.ExchangeDeclare(
		queueName,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		c.logger.Errorf("Error with ExchangeDeclare to RabbitMQ: ", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = ch.PublishWithContext(ctx, "", queueName, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(data),
	})
	if err != nil {
		c.logger.Errorf("Error with Publishing RabbitMQ: ", err)
	}
	return nil
}
