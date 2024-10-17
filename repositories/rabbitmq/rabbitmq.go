//go:generate mockgen -source=./rabbitmq.go -destination=../../mocks/rabbitmq_mock.go -package=mock
package rabbitmq

import (
	"card-project/config"
	"card-project/models"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgconn"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	headersMethod = "method"
	headersItem   = "item"
	rabbitConfigString = "amqp://%s:%s@%s:%s/"
)

type rabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel

	userRepo userRepo
	cardRepo cardRepo
}

type userRepo interface {
	PostUser(ctx context.Context, userData models.User) (models.User, error)
	DeleteUserID(ctx context.Context, id int) (pgconn.CommandTag, error)
}

type cardRepo interface {
	PostCard(ctx context.Context, cardData models.Card) (models.Card, error)
	DeleteCardID(ctx context.Context, id int) (pgconn.CommandTag, error)
}

type RabbitMQ interface {
	ProducePostUser(ctx context.Context, userData models.User) error
	ProduceDeleteUser(ctx context.Context, id int) error
	ProducePostCard(ctx context.Context, cardData models.Card) error
	ProduceDeleteCard(ctx context.Context, id int) error

	NewConsumer(ctx context.Context)
	ConsumeUserPost(ctx context.Context, msg amqp.Delivery)
	ConsumeUserDelete(ctx context.Context, msg amqp.Delivery)
	ConsumeCardPost(ctx context.Context, msg amqp.Delivery)
	ConsumeCardDelete(ctx context.Context, msg amqp.Delivery)
}

func NewConn(userRepo userRepo, cardRepo cardRepo, config config.Config) RabbitMQ {
	connString := fmt.Sprintf(rabbitConfigString, config.RabbitMQ.User, config.RabbitMQ.Password, config.RabbitMQ.Host, config.RabbitMQ.Port)
	conn, err := amqp.Dial(connString)
	if err != nil {
		log.Fatalf("Unable to connect to rabbitmq: %v\n", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel in rabbitmq: %v\n", err)
	}

	return &rabbitMQ{
		conn:     conn,
		channel:  channel,
		userRepo: userRepo,
		cardRepo: cardRepo,
	}
}
