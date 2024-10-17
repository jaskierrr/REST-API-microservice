package database

import (
	"card-project/config"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type db struct {
	conn *pgx.Conn
}

type DB interface {
	NewConn(ctx context.Context, connConfigString string, config config.Config) DB
	GetConn() *pgx.Conn
}

func NewDB() DB {
	return &db{}
}

func (d *db) NewConn(ctx context.Context, connConfigString string, config config.Config) DB {
	connString := fmt.Sprintf(connConfigString, config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)

	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		log.Fatalf("Failed to ping the database: %v\n", err)
	}

	return &db{
		conn: conn,
	}
}

func (db *db) GetConn() *pgx.Conn {
	return db.conn
}
