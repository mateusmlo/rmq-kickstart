package config

import (
	"fmt"

	utils "github.com/mateusmlo/rabbitmq-hello-world/utils"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

// AMQPEnvs basic environment variables to connect to rabbitmq instance
type AMQPEnvs struct {
	Username string
	Password string
	Host     string
	Port     int
}

// ConnectAMQP connects to underlying rabbitmq instance
func ConnectAMQP() (*amqp.Channel, *amqp.Queue, *amqp.Connection) {
	amqpEnvs := AMQPEnvs{
		Username: viper.GetString("RABBITMQ_USERNAME"),
		Password: viper.GetString("RABBITMQ_PASSWORD"),
		Host:     viper.GetString("RABBITMQ_NODE_NAME"),
		Port:     5672,
	}

	amqpConnURI := fmt.Sprintf("amqp://%s:%s@%s:%d/", amqpEnvs.Username, amqpEnvs.Password, amqpEnvs.Host, amqpEnvs.Port)

	conn, err := amqp.Dial(amqpConnURI)
	utils.FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError(err, "Failed to declare a queue")

	return ch, &q, conn
}
