package service

import (
	"card-project/models"
	"card-project/rabbitmq"
	cards_repo "card-project/repositories/cards"
	users_repo "card-project/repositories/users"
	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

type service struct {
	userRepo users_repo.UsersRepo
	cardRepo cards_repo.CardsRepo

	rabbitMQ rabbitmq.RabbitMQ
}

type Service interface {
	GetUserID(ctx context.Context, id string) (models.User, error)
	PostUser(ctx context.Context, user models.NewUser) (models.User, error)
	DeleteUserID(ctx context.Context, id string) (pgconn.CommandTag, error)
	GetUsers(ctx context.Context) ([]*models.User, error)

	GetCardID(ctx context.Context, id string) (models.Card, error)
	PostCard(ctx context.Context, user models.NewCard) (models.Card, error)
	DeleteCardID(ctx context.Context, id string) (pgconn.CommandTag, error)
	GetCards(ctx context.Context) ([]*models.Card, error)
}

func New(userRepo users_repo.UsersRepo, cardRepo cards_repo.CardsRepo, rabbitmq rabbitmq.RabbitMQ) Service {
	return service{
		userRepo: userRepo,
		cardRepo: cardRepo,
		rabbitMQ: rabbitmq,
	}
}
