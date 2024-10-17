//go:generate mockgen -source=./cards.go -destination=../../mocks/cards_repo_mock.go -package=mock
package cards_repo

import (
	"card-project/database"
	"card-project/models"
	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

const (
	getCardIDQuery    = `select * from cards where id = @cardID`
	postCardQuery     = `insert into cards (id, userid, bankid, number, created_at) values (@id, @userID, @bankID, @number, @create_date) returning *`
	deleteCardIDQuery = `delete from cards where id = @cardID`
	getCardsQuery     = `select * from cards`
)

type cardRepo struct {
	db database.DB
}

type CardsRepo interface {
	GetCardID(ctx context.Context, id int) (models.Card, error)
	PostCard(ctx context.Context, card models.Card) (models.Card, error)
	DeleteCardID(ctx context.Context, id int) (pgconn.CommandTag, error)
	GetCards(ctx context.Context) ([]*models.Card, error)
}

func NewCardRepo(db database.DB) CardsRepo {
	return &cardRepo{
		db: db,
	}
}
