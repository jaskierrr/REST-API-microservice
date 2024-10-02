package rabbitmq

import (
	"card-project/models"
	"context"
	"encoding/json"
	"log"
	"strconv"
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
			case msg.Headers["method"] == "POST":
				{
					userData := models.NewUser{}
					err := json.Unmarshal(msg.Body, &userData)
					if err != nil {
						log.Printf("Error unmarshaling message: %v", err)
						continue
					}

					_, err = r.userRepo.PostUser(ctx, userData)
					if err != nil {
						log.Printf("Error post user from consumer: %v", err)
						continue
					}

				}
			case msg.Headers["method"] == "DELETE":
				{
					body, _ := strconv.Atoi(string(msg.Body))

					_, err = r.userRepo.DeleteUserID(ctx, body)
					if err != nil {
						log.Printf("Error delete user from consumer: %v", err)
						continue
					}
					// fmt.Println(msg.Headers["method"])
					// fmt.Println(string(msg.Body))
				}
			}
		}
	}()

	select {}
}
