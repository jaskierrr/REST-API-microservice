package users_repo

import (
	"card-project/models"
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

func (repo *userRepo) PostUser(ctx context.Context, userData models.User) (models.User, error) {
	args := pgx.NamedArgs{
		"id":        userData.ID,
		"firstName": userData.FirstName,
		"lastName":  userData.LastName,
	}
	user := models.User{}
	err := repo.db.
		GetConn().
		QueryRow(ctx, postUserQuery, args).
		Scan(&user.ID, &user.FirstName, &user.LastName)

	if err != nil {
		return models.User{}, err
	}

	repo.logger.Info(
		"Success POST user from storage",
		slog.Any("ID", user.ID),
	)

	return user, nil
}
