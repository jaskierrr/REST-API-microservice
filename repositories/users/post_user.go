package users_repo

import (
	"card-project/models"
	"context"

	"github.com/jackc/pgx/v5"
)

func (repo *userRepo) PostUser(ctx context.Context, userData models.NewUser) (models.User, error) {
	args := pgx.NamedArgs{
		"firstName": userData.FirstName,
		"lastName":  userData.LastName,
	}
	user := models.User{}
	err := repo.db.
		GetConn().
		QueryRow(ctx, postUserQuery, args).
		Scan(&user.ID, &user.FirstName, &user.LastName)

	return user, err
}
