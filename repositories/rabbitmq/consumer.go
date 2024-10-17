package rabbitmq

import (
	"context"
	"log"
)

func (r *rabbitMQ) NewConsumer(ctx context.Context) {
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

	msgs, err := r.channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		log.Fatalf("Failed to declare a consumer in rabbitmq: %v\n", err)
	}

	go func() {
		for msg := range msgs {
			switch {
			case msg.Headers[headersItem] == "user":
				{
					switch {
					case msg.Headers[headersMethod] == "POST":
						{
							r.ConsumeUserPost(ctx, msg)
						}
					case msg.Headers[headersMethod] == "DELETE":
						{
							r.ConsumeUserDelete(ctx, msg)
						}
					}
				}
			case msg.Headers[headersItem] == "card":
				{
					switch {
					case msg.Headers[headersMethod] == "POST":
						{
							r.ConsumeCardPost(ctx, msg)
						}
					case msg.Headers[headersMethod] == "DELETE":
						{
							r.ConsumeCardDelete(ctx, msg)
						}
					}
				}
			}
		}
	}()

	select {}
}
