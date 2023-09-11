package worker

import (
	"bytes"
	"log"
	"time"

	utils "github.com/mateusmlo/rabbitmq-hello-world/tools"
	amqp "github.com/rabbitmq/amqp091-go"
)

// ProcessTask processes messages that takes a while
func ProcessTask(ch *amqp.Channel, q *amqp.Queue) {
	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError(err, "Failed to register a worker consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)

			time.Sleep(t * time.Second)
			log.Printf("Done processing")

			d.Ack(false)
		}
	}()

	log.Printf("[*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
