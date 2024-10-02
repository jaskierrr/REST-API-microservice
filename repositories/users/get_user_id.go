package users_repo

import (
	"card-project/models"
	"context"

	"github.com/jackc/pgx/v5"
)

func (repo *userRepo) GetUserID(ctx context.Context, id int) (models.User, error) {
	args := pgx.NamedArgs{
		"userID": id,
	}
	user := models.User{}
	err := repo.db.
		GetConn().
		QueryRow(ctx, getUserIDQuery, args).
		Scan(&user.ID, &user.FirstName, &user.LastName)

	return user, err
}
