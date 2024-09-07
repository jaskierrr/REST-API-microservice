package controller

import (
	"card-project/database"
	"card-project/models"
	"context"

	"github.com/jackc/pgx/v5"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "098098"
	dbname   = "CardProject"

	connConfigString = "postgres://%s:%s@%s:%d/%s"
)

type controller struct {
	db  database.DB
	ctx context.Context
}

type Controller interface {
	GetUserID(id string) (models.NewUser, error)
}

func New() Controller {
	return controller{
		db:  database.NewDB().NewConn(connConfigString, user, password, host, port, dbname),
		ctx: context.Background(),
	}
}

func (c controller) GetUserID(id string) (models.NewUser, error) {
	query := `select first_name, last_name from users where id = @userID`

	args := pgx.NamedArgs{
		"userID": id,
	}

	user := models.NewUser{}

	err := c.db.GetConn().QueryRow(c.ctx, query, args).Scan(&user.FirstName, &user.LastName)

	return user, err
}
