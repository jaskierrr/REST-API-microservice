package service

import (
	"card-project/models"
	"context"

	"github.com/google/uuid"
)

func (s service) GetUserID(ctx context.Context, id int) (models.User, error) {
	return s.userRepo.GetUserID(ctx, id)
}

func (s service) PostUser(ctx context.Context, userData models.NewUser) (models.User, error) {
	id, _ := uuid.NewUUID()

	user := models.User{
		ID:        int64(id.ID()),
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
	}
	err := s.rabbitMQ.ProducePostUser(ctx, user)

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s service) DeleteUserID(ctx context.Context, id int) error {
	return s.rabbitMQ.ProduceDeleteUser(ctx, id)
}

func (s service) GetUsers(ctx context.Context) ([]*models.User, error) {
	return s.userRepo.GetUsers(ctx)
}
