package rabbitmq

import (
	"card-project/models"
	"context"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (r *rabbitMQ) ProducePostCard(ctx context.Context, userData models.NewCard) {
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

	body, err := json.Marshal(userData)
	if err != nil {
		log.Fatalf("Failed to marshal userData in json: %v\n", err)
	}

	headers := make(amqp.Table)
	headers["method"] = "POST"

	err = r.channel.PublishWithContext(
		ctx,
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			Headers:     headers,
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish data in rabbitmq: %v\n", err)
	}

}
