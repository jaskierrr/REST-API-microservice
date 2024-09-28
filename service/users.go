package service

import (
	"card-project/models"
	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

func (s service) GetUserID(ctx context.Context, id string) (models.User, error) {
	user, err := s.userRepo.GetUserID(ctx, id)

	return user, err
}

func (s service) PostUser(ctx context.Context, userData models.NewUser) (models.User, error) {
	s.rabbitMQ.ProduceUsersPOST(ctx, userData)

	user, err := s.userRepo.PostUser(ctx, userData)


	return user, err
}

func (s service) DeleteUserID(ctx context.Context, id string) (pgconn.CommandTag, error) {
	s.rabbitMQ.ProduceUsersDELETE(ctx, id)

	commandTag, err := s.userRepo.DeleteUserID(ctx, id)

	return commandTag, err
}

func (s service) GetUsers(ctx context.Context) ([]*models.User, error) {
	user, err := s.userRepo.GetUsers(ctx)

	return user, err
}
