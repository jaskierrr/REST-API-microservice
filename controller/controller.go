package controller

import (
	"card-project/models"
	"card-project/service"
	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

type controller struct {
	service service.Service
}

type Controller interface {
	GetUserID(ctx context.Context, id string) (models.User, error)
	PostUser(ctx context.Context, user models.NewUser) (models.User, error)
	DeleteUserID(ctx context.Context, id string) (pgconn.CommandTag, error)
	GetUsers(ctx context.Context) ([]*models.User, error)
}

func New(service service.Service) Controller {
	return controller{
		service: service,
	}
}

func (c controller) GetUserID(ctx context.Context, id string) (models.User, error) {
	user, err := c.service.GetUserID(ctx, id)

	return user, err
}

func (c controller) PostUser(ctx context.Context, userData models.NewUser) (models.User, error) {
	user, err := c.service.PostUser(ctx, userData)

	return user, err
}

func (c controller) DeleteUserID(ctx context.Context, id string) (pgconn.CommandTag, error) {
	commandTag, err := c.service.DeleteUserID(ctx, id)

	return commandTag, err
}

func (c controller) GetUsers(ctx context.Context) ([]*models.User, error) {
	user, err := c.service.GetUsers(ctx)

	return user, err
}
