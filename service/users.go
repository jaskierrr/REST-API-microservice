package service

import (
	"card-project/models"
	"context"

	"github.com/google/uuid"
)

func (s service) GetUserID(ctx context.Context, id int) (models.User, error) {
	user, err := s.userRepo.GetUserID(ctx, id)

	return user, err
}

func (s service) PostUser(ctx context.Context, userData models.NewUser) (models.User, error) {
	id, _ := uuid.NewUUID()

	user := models.User{
		ID:        int64(id.ID()),
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
	}

	err := s.rabbitMQ.ProducePostUser(ctx, user)

	return user, err
}

func (s service) DeleteUserID(ctx context.Context, id int) error {
	err := s.rabbitMQ.ProduceDeleteUser(ctx, id)

	return err
}

func (s service) GetUsers(ctx context.Context) ([]*models.User, error) {
	user, err := s.userRepo.GetUsers(ctx)

	return user, err
}
