package cards_repo

import (
	"card-project/models"
	"context"

	"github.com/jackc/pgx/v5"
)

func (repo *cardRepo) GetCardID(ctx context.Context, id string) (models.Card, error) {
	args := pgx.NamedArgs{
		"cardID": id,
	}
	card := models.Card{}
	err := repo.db.GetConn().QueryRow(ctx, getCardIDQuery, args).Scan(&card.ID, &card.UserID, &card.BankID, &card.Number, &card.CreateDate)

	return card, err
}
