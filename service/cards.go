package service

import (
	"card-project/models"
	"context"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

func (s service) GetCardID(ctx context.Context, id int) (models.Card, error) {
	card, err := s.cardRepo.GetCardID(ctx, id)

	return card, err
}

func (s service) PostCard(ctx context.Context, cardData models.NewCard) (models.Card, error) {
	id, _ := uuid.NewUUID()

	card := models.Card{
		ID:         int64(id.ID()),
		UserID:     cardData.UserID,
		BankID:     cardData.BankID,
		Number:     cardData.Number,
		CreateDate: strfmt.DateTime(time.Now()),
	}

	err := s.rabbitMQ.ProducePostCard(ctx, card)

	return card, err
}

func (s service) DeleteCardID(ctx context.Context, id int) error {
	err := s.rabbitMQ.ProduceDeleteCard(ctx, id)

	return err
}

func (s service) GetCards(ctx context.Context) ([]*models.Card, error) {
	card, err := s.cardRepo.GetCards(ctx)

	return card, err
}
