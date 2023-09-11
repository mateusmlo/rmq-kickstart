package task

import (
	"context"
	"log"
	"os"
	"time"

	utils "github.com/mateusmlo/rabbitmq-hello-world/tools"
	amqp "github.com/rabbitmq/amqp091-go"
)

// CreateTask creates new message that takes a while to process
func CreateTask(ch *amqp.Channel, q *amqp.Queue) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := utils.BodyFrom(os.Args)

	err := ch.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(body),
	})

	utils.FailOnError(err, "Failed to publish message from new task")
	log.Printf("[x] Sent %s", body)
}
