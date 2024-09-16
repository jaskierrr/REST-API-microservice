package repositories

import (
	"card-project/database"
	"card-project/models"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	getCardIDQuery    = `select * from cards where id = @cardID`
	postCardQuery     = `insert into cards (number, userid, bankid, created_at) values (@number, @userID, @bankID, now()) returning *`
	deleteCardIDQuery = `delete from cards where id = @cardID`
	getCardsQuery     = `select * from cards`
)

type cardRepo struct {
	db database.DB
}

type CardsRepo interface {
	GetCardID(ctx context.Context, id string) (models.Card, error)
	PostCard(ctx context.Context, card models.NewCard) (models.Card, error)
	DeleteCardID(ctx context.Context, id string) (pgconn.CommandTag, error)
	GetCards(ctx context.Context) ([]*models.Card, error)
}

func NewCardRepo(db database.DB) CardsRepo {
	return &cardRepo{
		db: db,
	}
}

func (repo *cardRepo) GetCardID(ctx context.Context, id string) (models.Card, error) {
	args := pgx.NamedArgs{
		"cardID": id,
	}
	card := models.Card{}
	err := repo.db.GetConn().QueryRow(ctx, getCardIDQuery, args).Scan(&card.ID, &card.UserID, &card.BankID, &card.Number, &card.CreateDate)

	return card, err
}

func (repo *cardRepo) PostCard(ctx context.Context, cardData models.NewCard) (models.Card, error) {
	args := pgx.NamedArgs{
		"number": cardData.Number,
		"userID": cardData.UserID,
		"bankID": cardData.BankID,
	}
	card := models.Card{}
	err := repo.db.GetConn().QueryRow(ctx, postCardQuery, args).Scan(&card.ID, &card.UserID, &card.BankID, &card.Number, &card.CreateDate)

	return card, err
}

func (repo *cardRepo) DeleteCardID(ctx context.Context, id string) (pgconn.CommandTag, error) {
	args := pgx.NamedArgs{
		"cardID": id,
	}
	commandTag, err := repo.db.GetConn().Exec(ctx, deleteCardIDQuery, args)

	return commandTag, err
}

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
