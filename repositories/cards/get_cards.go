package cards_repo

import (
	"card-project/models"
	"context"
)

func (repo *cardRepo) GetCards(ctx context.Context) ([]*models.Card, error) {

	rows, err := repo.db.GetConn().Query(ctx, getCardsQuery)

	cards := []*models.Card{}
	defer rows.Close()

	for rows.Next() {
		card := models.Card{}

		err := rows.Scan(&card.ID, &card.UserID, &card.BankID, &card.Number, &card.CreateDate)
		if err != nil {
			return nil, err
		}

		cards = append(cards, &card)
	}

	return cards, err
}
