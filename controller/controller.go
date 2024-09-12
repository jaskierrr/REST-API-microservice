package controller

import (
	"card-project/models"
	repositories "card-project/repositories/users"
	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

type controller struct {
	repository repositories.UsersRepo
}

type Controller interface {
	GetUserID(ctx context.Context, id string) (models.User, error)
	PostUser(ctx context.Context, user models.NewUser) (models.User, error)
	DeleteUserID(ctx context.Context, id string) (pgconn.CommandTag, error)
	GetUsers(ctx context.Context) ([]*models.User, error)
}

func New(repository repositories.UsersRepo) Controller {
	return controller{
		repository: repository,
	}
}

func (c controller) GetUserID(ctx context.Context, id string) (models.User, error) {
	user, err := c.repository.GetUserID(ctx, id)

	return user, err
}

func (c controller) PostUser(ctx context.Context, userData models.NewUser) (models.User, error) {
	user, err := c.repository.PostUser(ctx, userData)

	return user, err
}

func (c controller) DeleteUserID(ctx context.Context, id string) (pgconn.CommandTag, error) {
	commandTag, err := c.repository.DeleteUserID(ctx, id)

	return commandTag, err
}

func (c controller) GetUsers(ctx context.Context) ([]*models.User, error) {
	user, err := c.repository.GetUsers(ctx)

	return user, err
}
