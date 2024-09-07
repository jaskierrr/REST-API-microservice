package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)


var ctx = context.Background()

type db struct {
	Conn *pgx.Conn
}

type DB interface {
	NewConn(connConfigString string, user string, password string, host string, port int, dbname string) DB
	GetConn() *pgx.Conn
}

func NewDB() DB {
	return &db{}
}


func (db *db) NewConn(connConfigString string, user string, password string, host string, port int, dbname string) DB{

	connString := fmt.Sprintf(connConfigString, user, password, host, port, dbname)

	conn, err := pgx.Connect(ctx, connString)
	// defer db.Close(context.Background())
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		log.Fatalf("Failed to ping the database: %v\n", err)
	}

	db.Conn = conn

	return db
}

func (db *db) GetConn() *pgx.Conn{
	return db.Conn
}
