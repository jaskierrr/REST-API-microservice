package cards_repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (repo *cardRepo) DeleteCardID(ctx context.Context, id string) (pgconn.CommandTag, error) {
	args := pgx.NamedArgs{
		"cardID": id,
	}
	commandTag, err := repo.db.
		GetConn().
		Exec(ctx, deleteCardIDQuery, args)

	return commandTag, err
}
