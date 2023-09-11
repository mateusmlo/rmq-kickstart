package logs

import (
	"context"
	"log"
	"os"
	"time"

	utils "github.com/mateusmlo/rabbitmq-hello-world/tools"
	amqp "github.com/rabbitmq/amqp091-go"
)

// EmitLogs publish message to exchange
func EmitLogs(ch *amqp.Channel) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := utils.BodyFrom(os.Args)

	err := ch.PublishWithContext(ctx, "logs", "", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	utils.FailOnError(err, "Failed to publish to exchange")

	log.Printf("[x] Send %s", body)
}
