package cards_repo

import (
	"card-project/models"
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

func (repo *cardRepo) GetCardID(ctx context.Context, id int) (models.Card, error) {
	args := pgx.NamedArgs{
		"cardID": id,
	}
	card := models.Card{}
	err := repo.db.
		GetConn().
		QueryRow(ctx, getCardIDQuery, args).
		Scan(&card.ID, &card.UserID, &card.BankID, &card.Number, &card.CreateDate)

	if err != nil {
		return models.Card{}, err
	}

	repo.logger.Info(
		"Success GET card from storage",
		slog.Any("ID", card.ID),
	)


	return card, nil
}
