package rabbitmq

import (
	"card-project/models"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgconn"
	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel

	userRepo userRepo
}

type userRepo interface {
	PostUser(ctx context.Context, userData models.NewUser) (models.User, error)
	DeleteUserID(ctx context.Context, id int) (pgconn.CommandTag, error)
}

type RabbitMQ interface {
	NewConn(userRepo userRepo) RabbitMQ
	ProduceUsersPOST(ctx context.Context, userData models.NewUser)
	ProduceUsersDELETE(ctx context.Context, id int)
	NewConsumer(ctx context.Context)
}

func NewRabbitMQ() RabbitMQ {
	return &rabbitMQ{}
}

func (r *rabbitMQ) NewConn(userRepo userRepo) RabbitMQ {

	conn, err := amqp.Dial("amqp://jaskier:test@rabbitmq:5672/")
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
	}

}
