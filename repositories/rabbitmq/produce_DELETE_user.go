package rabbitmq

import (
	"context"
	"log"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (r *rabbitMQ) ProduceUsersDELETE(ctx context.Context, id int) {
	queue, err := r.channel.QueueDeclare(
		"test", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue in rabbitmq: %v\n", err)
	}

	headers := make(amqp.Table)
	headers["method"] = "DELETE"

	err = r.channel.PublishWithContext(
		ctx,
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			Headers:     headers,
			ContentType: "text/plain",
			Body:        []byte(strconv.Itoa(id)),
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish data in rabbitmq: %v\n", err)
	}
}
