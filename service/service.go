package service

import (
	"card-project/models"
	repositories "card-project/repositories/users"
	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

type service struct {
	repository repositories.UsersRepo
}

type Service interface {
	GetUserID(ctx context.Context, id string) (models.User, error)
	PostUser(ctx context.Context, user models.NewUser) (models.User, error)
	DeleteUserID(ctx context.Context, id string) (pgconn.CommandTag, error)
	GetUsers(ctx context.Context) ([]*models.User, error)
}

func New(repository repositories.UsersRepo) Service {
	return service{
		repository: repository,
	}
}

func (s service) GetUserID(ctx context.Context, id string) (models.User, error) {
	user, err := s.repository.GetUserID(ctx, id)

	return user, err
}

func (s service) PostUser(ctx context.Context, userData models.NewUser) (models.User, error) {
	user, err := s.repository.PostUser(ctx, userData)

	return user, err
}

func (s service) DeleteUserID(ctx context.Context, id string) (pgconn.CommandTag, error) {
	commandTag, err := s.repository.DeleteUserID(ctx, id)

	return commandTag, err
}

func (s service) GetUsers(ctx context.Context) ([]*models.User, error) {
	user, err := s.repository.GetUsers(ctx)

	return user, err
}
