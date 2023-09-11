package logs

import (
	"log"

	utils "github.com/mateusmlo/rabbitmq-hello-world/tools"
	amqp "github.com/rabbitmq/amqp091-go"
)

// ReceiveLogs receives logs published to exchange
func ReceiveLogs(ch *amqp.Channel) {
	err := ch.QueueBind("", "", "logs", false, nil)
	utils.FailOnError(err, "Failed to bind to queue")

	msgs, err := ch.Consume(
		"",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError(err, "Failed to register logs consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("[x] %s", d.Body)
		}
	}()

	log.Printf("[*] Waiting for logs. To exist press CTRL+C")
	<-forever
}
