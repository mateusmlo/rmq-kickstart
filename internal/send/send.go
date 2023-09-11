package send

import (
	"context"
	"log"
	"time"

	utils "github.com/mateusmlo/rabbitmq-hello-world/tools"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Send sends message to queue
func Send(ch *amqp.Channel, queue *amqp.Queue) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	err := ch.PublishWithContext(ctx, "", queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})

	utils.FailOnError(err, "Failed to publish message")
	log.Printf("[x] Sent %s\n", body)
}
