//go:generate mockgen -source=./users.go -destination=../../mocks/users_repo_mock.go -package=mock
package users_repo

import (
	"card-project/database"
	"card-project/models"
	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

const (
	getUserIDQuery    = `select * from users where id = @userID`
	postUserQuery     = `insert into users (id ,first_name, last_name) values (@id, @firstName, @lastName) returning *`
	deleteUserIDQuery = `delete from users where id = @userID`
	getUsersQuery     = `select * from users`
)

type userRepo struct {
	db database.DB
}

type UsersRepo interface {
	GetUserID(ctx context.Context, id int) (models.User, error)
	PostUser(ctx context.Context, user models.User) (models.User, error)
	DeleteUserID(ctx context.Context, id int) (pgconn.CommandTag, error)
	GetUsers(ctx context.Context) ([]*models.User, error)
}

func NewUserRepo(db database.DB) UsersRepo {
	return &userRepo{
		db: db,
	}
}
