package controller

import (
	"card-project/models"
	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

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
