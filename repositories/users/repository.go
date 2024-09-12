package repositories

import (
	"card-project/database"
	"card-project/models"
	"context"

	"github.com/jackc/pgx/v5"
)

const (
	getUserIDQuery = `select * from users where id = @userID`
	postUserQuery = `insert into users (first_name, last_name) values (@firstName, @lastName) returning *`
	deleteUserIDQuery = `delete from users where id = @userID`
)

type userRepo struct {
	db database.DB
}

type UsersRepo interface {
	GetUserID(ctx context.Context, id string) (models.User, error)
	PostUser(ctx context.Context, user models.NewUser) (models.User, error)
	DeleteUserID(ctx context.Context, id string) error
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
		"lastName": userData.LastName,
	}
	user := models.User{}
	err :=  repo.db.GetConn().QueryRow(ctx, postUserQuery, args).Scan(&user.ID, &user.FirstName, &user.LastName)

	return user, err
}

func (repo *userRepo) DeleteUserID(ctx context.Context, id string) error {
	args := pgx.NamedArgs{
		"userID": id,
	}
	err := repo.db.GetConn().Query()
	return err
}
