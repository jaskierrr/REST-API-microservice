package controller

import (
	"card-project/models"
	"context"

)

func (c controller) GetUserID(ctx context.Context, id int) (models.User, error) {
	user, err := c.service.GetUserID(ctx, id)
	return user, err
}

func (c controller) PostUser(ctx context.Context, userData models.NewUser) (models.User, error) {
	user, err := c.service.PostUser(ctx, userData)
	return user, err
}

func (c controller) DeleteUserID(ctx context.Context, id int) error {
	err := c.service.DeleteUserID(ctx, id)
	return err
}

func (c controller) GetUsers(ctx context.Context) ([]*models.User, error) {
	user, err := c.service.GetUsers(ctx)
	return user, err
}
