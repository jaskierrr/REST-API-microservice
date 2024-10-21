package users_repo

import (
	"card-project/models"
	"context"

)

func (repo *userRepo) GetUsers(ctx context.Context) ([]*models.User, error) {

	rows, err := repo.db.GetConn().Query(ctx, getUsersQuery)

	users := []*models.User{}
	defer rows.Close()

	for rows.Next() {
		user := models.User{}

		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	repo.logger.Info("Success GET users from storage")

	return users, err
}
