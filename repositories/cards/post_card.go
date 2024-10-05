package cards_repo

import (
	"card-project/models"
	"context"

	"github.com/jackc/pgx/v5"
)

func (repo *cardRepo) PostCard(ctx context.Context, cardData models.NewCard) (models.Card, error) {
	args := pgx.NamedArgs{
		"number": cardData.Number,
		"userID": cardData.UserID,
		"bankID": cardData.BankID,
	}
	card := models.Card{}
	err := repo.db.
		GetConn().
		QueryRow(ctx, postCardQuery, args).
		Scan(&card.ID, &card.UserID, &card.BankID, &card.CreateDate, &card.Number)

	return card, err
}
