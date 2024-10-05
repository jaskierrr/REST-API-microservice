package service

import (
	"card-project/models"
	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

func (s service) GetUserID(ctx context.Context, id int) (models.User, error) {
	user, err := s.userRepo.GetUserID(ctx, id)

	return user, err
}

func (s service) PostUser(ctx context.Context, userData models.NewUser) (models.User, error) {
	s.rabbitMQ.ProducePostUser(ctx, userData)

	// user, err := s.userRepo.PostUser(ctx, userData)

	return models.User{}, nil
}

func (s service) DeleteUserID(ctx context.Context, id int) (pgconn.CommandTag, error) {
	s.rabbitMQ.ProduceDeleteUser(ctx, id)

	// commandTag, err := s.userRepo.DeleteUserID(ctx, id)

	return pgconn.CommandTag{}, nil
}

func (s service) GetUsers(ctx context.Context) ([]*models.User, error) {
	user, err := s.userRepo.GetUsers(ctx)

	return user, err
}
