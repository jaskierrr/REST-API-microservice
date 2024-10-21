package cards_repo

import (
	"card-project/models"
	"context"
	"log/slog"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/jackc/pgx/v5"
)

func (repo *cardRepo) PostCard(ctx context.Context, cardData models.Card) (models.Card, error) {
	args := pgx.NamedArgs{
		"id":          cardData.ID,
		"userID":      cardData.UserID,
		"bankID":      cardData.BankID,
		"number":      cardData.Number,
		"create_date": time.Time(cardData.CreateDate),
	}
	var createTime time.Time
	card := models.Card{}
	err := repo.db.
		GetConn().
		QueryRow(ctx, postCardQuery, args).
		Scan(&card.ID, &card.UserID, &card.BankID, &card.Number, &createTime)

	card.CreateDate = strfmt.DateTime(createTime)

	if err != nil {
		return models.Card{}, err
	}

	repo.logger.Info(
		"Success POST card from storage",
		slog.Any("ID", card.ID),
	)

	return card, nil
}
