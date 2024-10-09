package rabbitmq

import (
	"card-project/models"
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (r *rabbitMQ) ProducePostCard(ctx context.Context, cardData models.Card) error {
	queue, err := r.channel.QueueDeclare(
		"test", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		// log.Printf("Failed to declare a queue in rabbitmq: %v\n", err)
		return err
	}

	body, err := json.Marshal(cardData)
	if err != nil {
		// log.Printf("Failed to marshal userData in json: %v\n", err)
		return err
	}

	headers := make(amqp.Table)
	headers[headersMethod] = "POST"
	headers[headersItem] = "card"

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
		// log.Printf("Failed to publish data in rabbitmq: %v\n", err)
		return err
	}

	return nil
}
