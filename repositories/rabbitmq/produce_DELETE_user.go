package rabbitmq

import (
	"context"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (r *rabbitMQ) ProduceDeleteUser(ctx context.Context, id int) error {
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

	headers := make(amqp.Table)
	headers[headersMethod] = "DELETE"
	headers[headersItem] = "user"

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
		// log.Printf("Failed to publish data in rabbitmq: %v\n", err)
		return err
	}

	return nil
}
