package config

import (
	"fmt"

	utils "github.com/mateusmlo/rabbitmq-hello-world/tools"
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

func newAMQPCredentials(username, password, host string) AMQPEnvs {
	return AMQPEnvs{
		Username: username,
		Password: password,
		Host:     host,
		Port:     5672,
	}
}

// Connect connects to underlying rabbitmq instance
func Connect() (*amqp.Channel, *amqp.Connection) {
	amqpEnvs := newAMQPCredentials(
		viper.GetString("RABBITMQ_USERNAME"),
		viper.GetString("RABBITMQ_PASSWORD"),
		viper.GetString("RABBITMQ_NODE_NAME"),
	)

	amqpConnURI := fmt.Sprintf("amqp://%s:%s@%s:%d/", amqpEnvs.Username, amqpEnvs.Password, amqpEnvs.Host, amqpEnvs.Port)

	conn, err := amqp.Dial(amqpConnURI)
	utils.FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")

	err = ch.Qos(2, 0, false)

	return ch, conn
}

// CreateQueue declares a new queue
func CreateQueue(ch *amqp.Channel, qName string, durable bool, exclusive bool) (q amqp.Queue, err error) {
	q, err = ch.QueueDeclare(
		qName,
		durable,
		false,
		exclusive,
		false,
		nil,
	)
	utils.FailOnError(err, "Failed to declare a queue")

	return
}

// CreateExchange declares a new exchange
func CreateExchange(ch *amqp.Channel, exName string) (err error) {
	err = ch.ExchangeDeclare(
		exName,
		amqp.ExchangeFanout,
		true,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError(err, "Failed to declare an exchange")

	return
}
