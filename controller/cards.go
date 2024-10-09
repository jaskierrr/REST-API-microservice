package controller

import (
	"card-project/models"
	"context"
)

func (c controller) GetCardID(ctx context.Context, id int) (models.Card, error) {
	card, err := c.service.GetCardID(ctx, id)
	return card, err
}

func (c controller) PostCard(ctx context.Context, cardData models.NewCard) (models.Card, error) {
	card, err := c.service.PostCard(ctx, cardData)
	return card, err
}

func (c controller) DeleteCardID(ctx context.Context, id int) error {
	err := c.service.DeleteCardID(ctx, id)
	return err
}

func (c controller) GetCards(ctx context.Context) ([]*models.Card, error) {
	card, err := c.service.GetCards(ctx)
	return card, err
}
