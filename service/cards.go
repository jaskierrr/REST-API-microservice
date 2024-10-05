package service

import (
	"card-project/models"
	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

func (s service) GetCardID(ctx context.Context, id int) (models.Card, error) {
	card, err := s.cardRepo.GetCardID(ctx, id)

	return card, err
}

func (s service) PostCard(ctx context.Context, cardData models.NewCard) (models.Card, error) {
	s.rabbitMQ.ProducePostCard(ctx, cardData)

	// card, err := s.cardRepo.PostCard(ctx, cardData)

	return models.Card{}, nil
}

func (s service) DeleteCardID(ctx context.Context, id int) (pgconn.CommandTag, error) {
	s.rabbitMQ.ProduceDeleteCard(ctx, id)

	// commandTag, err := s.cardRepo.DeleteCardID(ctx, id)

	return pgconn.CommandTag{}, nil
}

func (s service) GetCards(ctx context.Context) ([]*models.Card, error) {
	card, err := s.cardRepo.GetCards(ctx)

	return card, err
}
