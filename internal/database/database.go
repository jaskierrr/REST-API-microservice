package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "098098"
	dbname   = "CardProject"
)

var ctx = context.Background()

type postgres struct {
	Conn *pgx.Conn
}

var postgresInstance postgres = postgres{}

func OpenConn() *postgres{
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, dbname)

	conn, err := pgx.Connect(ctx, connString)
	// defer db.Close(context.Background())
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		log.Fatalf("Failed to ping the database: %v\n", err)
	}

	postgresInstance.Conn = conn

	return &postgresInstance
}


func AddUser(firstName string, lastName string) {
	query := `insert into users (first_name, last_name) values (@firstName, @lastName)`

	args := pgx.NamedArgs{
		"firstName": firstName,
		"lastName": lastName,
	}

	_, err := postgresInstance.Conn.Exec(ctx, query, args)
	if err != nil {
		log.Fatal("Failed to add User in storage: ", err)
	}

}

func DeleteUser(id string) {
	query := `delete from users where id = @userID`

	args := pgx.NamedArgs{
		"userID": id,
	}

	_, err := postgresInstance.Conn.Exec(ctx, query, args)
	if err != nil {
		log.Fatal("Failed to delete User in storage: ", err)
	}

}

func UpdateUser(id string, firstName string, lastName string) {
	//UPDATE users SET name = $1, email = $2 WHERE id = $3
	query := `update users set first_name = @firstName, last_name = @lastName where id = @userID`

	args := pgx.NamedArgs{
		"userID": id,
		"firstName": firstName,
		"lastName": lastName,
	}

	_, err := postgresInstance.Conn.Exec(ctx, query, args)
	if err != nil {
		log.Fatal("Failed to update User in storage: ", err)
	}

}
