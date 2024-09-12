package controller

import (
	"card-project/models"
	repositories "card-project/repositories/users"
	"context"
)

type controller struct {
	repository repositories.UsersRepo
}

type Controller interface {
	GetUserID(ctx context.Context, id string) (models.User, error)
	PostUser(ctx context.Context, user models.NewUser) (models.User, error)
	DeleteUserID(ctx context.Context, id string) error
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

func (c controller) DeleteUserID(ctx context.Context, id string) error {
	err := c.repository.DeleteUserID(ctx, id)

	return err
}
