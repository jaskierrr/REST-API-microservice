package users_repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (repo *userRepo) DeleteUserID(ctx context.Context, id string) (pgconn.CommandTag, error) {
	args := pgx.NamedArgs{
		"userID": id,
	}
	commandTag, err := repo.db.
		GetConn().
		Exec(ctx, deleteUserIDQuery, args)

	return commandTag, err
}
