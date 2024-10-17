package rabbitmq

import (
	"card-project/models"
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/rabbitmq/amqp091-go"
)

func (r *rabbitMQ) consumeUserPost(ctx context.Context, msg amqp091.Delivery) {
	userData := models.User{}
	err := json.Unmarshal(msg.Body, &userData)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
	}
	_, err = r.userRepo.PostUser(ctx, userData)
	if err != nil {
		log.Printf("Error post user from consumer: %v", err)
	}
}

func (r *rabbitMQ) consumeUserDelete(ctx context.Context, msg amqp091.Delivery) {
	body, _ := strconv.Atoi(string(msg.Body))
	commandTag, err := r.userRepo.DeleteUserID(ctx, body)
	if commandTag.RowsAffected() == 0 {
		log.Printf("Error delete user from consumer: %v", err)
	}
	if err != nil {
		log.Printf("Error delete user from consumer: %v", err)
	}
}
