//go:generate mockgen -source=./service.go -destination=../mocks/service_mock.go -package=mock
package service

import (
	"card-project/models"
	cards_repo "card-project/repositories/cards"
	"card-project/repositories/rabbitmq"
	users_repo "card-project/repositories/users"
	"context"
	"log/slog"
)

type service struct {
	logger *slog.Logger

	userRepo users_repo.UsersRepo
	cardRepo cards_repo.CardsRepo
	rabbitMQ rabbitmq.RabbitMQ
}

type Service interface {
	GetUserID(ctx context.Context, id int) (models.User, error)
	PostUser(ctx context.Context, user models.NewUser) (models.User, error)
	DeleteUserID(ctx context.Context, id int) error
	GetUsers(ctx context.Context) ([]*models.User, error)

	GetCardID(ctx context.Context, id int) (models.Card, error)
	PostCard(ctx context.Context, user models.NewCard) (models.Card, error)
	DeleteCardID(ctx context.Context, id int) error
	GetCards(ctx context.Context) ([]*models.Card, error)
}

func New(userRepo users_repo.UsersRepo, cardRepo cards_repo.CardsRepo, rabbitmq rabbitmq.RabbitMQ, logger *slog.Logger) Service {
	return &service{
		logger: logger,

		userRepo: userRepo,
		cardRepo: cardRepo,
		rabbitMQ: rabbitmq,
	}
}
