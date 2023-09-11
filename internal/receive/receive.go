package receive

import (
	"log"

	utils "github.com/mateusmlo/rabbitmq-hello-world/tools"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Receive consumes messages from queue
func Receive(ch *amqp.Channel, queue *amqp.Queue) {

	msgs, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError(err, "Failed to register consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received message: %s\n", d.Body)
		}
	}()

	log.Printf("[*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
