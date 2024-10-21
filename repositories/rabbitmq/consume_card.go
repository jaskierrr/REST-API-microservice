package rabbitmq

import (
	"card-project/models"
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/rabbitmq/amqp091-go"
)

func (r *rabbitMQ) ConsumeCardPost(ctx context.Context, msg amqp091.Delivery) {
	cardData := models.Card{}
	err := json.Unmarshal(msg.Body, &cardData)
	if err != nil {
		log.Printf("Error unmarshaling message: %v", err)
	}
	_, err = r.cardRepo.PostCard(ctx, cardData)
	if err != nil {
		log.Printf("Error post card from consumer: %v", err)
	}
}
func (r *rabbitMQ) ConsumeCardDelete(ctx context.Context, msg amqp091.Delivery) {
	body, _ := strconv.Atoi(string(msg.Body))
	commandTag, err := r.cardRepo.DeleteCardID(ctx, body)
	if commandTag.RowsAffected() == 0 {
		log.Printf("Error delete card from consumer: %v", err)
	}
	if err != nil {
		log.Printf("Error delete card from consumer: %v", err)
	}
}
