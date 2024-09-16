package repositories

import (
	"card-project/database"
	"card-project/models"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	getUserIDQuery    = `select * from users where id = @userID`
	postUserQuery     = `insert into users (first_name, last_name) values (@firstName, @lastName) returning *`
	deleteUserIDQuery = `delete from users where id = @userID`
	getUsersQuery     = `select * from users`
)

type userRepo struct {
	db database.DB
}

type UsersRepo interface {
	GetUserID(ctx context.Context, id string) (models.User, error)
	PostUser(ctx context.Context, user models.NewUser) (models.User, error)
	DeleteUserID(ctx context.Context, id string) (pgconn.CommandTag, error)
	GetUsers(ctx context.Context) ([]*models.User, error)
}

func NewUserRepo(db database.DB) UsersRepo {
	return &userRepo{
		db: db,
	}
}

func (repo *userRepo) GetUserID(ctx context.Context, id string) (models.User, error) {
	args := pgx.NamedArgs{
		"userID": id,
	}
	user := models.User{}
	err := repo.db.GetConn().QueryRow(ctx, getUserIDQuery, args).Scan(&user.ID, &user.FirstName, &user.LastName)

	return user, err
}

func (repo *userRepo) PostUser(ctx context.Context, userData models.NewUser) (models.User, error) {
	args := pgx.NamedArgs{
		"firstName": userData.FirstName,
		"lastName":  userData.LastName,
	}
	user := models.User{}
	err := repo.db.GetConn().QueryRow(ctx, postUserQuery, args).Scan(&user.ID, &user.FirstName, &user.LastName)

	return user, err
}

func (repo *userRepo) DeleteUserID(ctx context.Context, id string) (pgconn.CommandTag, error) {
	args := pgx.NamedArgs{
		"userID": id,
	}
	commandTag, err := repo.db.GetConn().Exec(ctx, deleteUserIDQuery, args)

	return commandTag, err
}

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

	return users, err
}
