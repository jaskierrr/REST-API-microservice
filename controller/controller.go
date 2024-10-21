//go:generate mockgen -source=./controller.go -destination=../mocks/controller_mock.go -package=mock
package controller

import (
	"card-project/models"
	"card-project/service"
	"context"
	"log/slog"
)

type controller struct {
	logger *slog.Logger
	service service.Service
}

type Controller interface {
	GetUserID(ctx context.Context, id int) (models.User, error)
	PostUser(ctx context.Context, user models.NewUser) (models.User, error)
	DeleteUserID(ctx context.Context, id int) error
	GetUsers(ctx context.Context) ([]*models.User, error)

	GetCardID(ctx context.Context, id int) (models.Card, error)
	PostCard(ctx context.Context, user models.NewCard) (models.Card, error)
	DeleteCardID(ctx context.Context, id int) error
	GetCards(ctx context.Context) ([]*models.Card, error)
}

func New(service service.Service, logger *slog.Logger) Controller {
	return &controller{
		logger: logger,
		service: service,
	}
}
