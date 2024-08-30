package database

import "log"

func CreateUsersTable(db *postgres) {
	query := `create table if not exists users (
		id serial primary key,
		first_name varchar(50) not null,
		last_name varchar(50) not null,
		created_at timestamp
	)`

	_, err := db.Conn.Exec(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

}

func CreateBanksTable(db *postgres) {
	query := `create table if not exists banks (
		id serial primary key,
		name varchar(50) not null,
		created_at timestamp
	)`

	_, err := db.Conn.Exec(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateCardsTable(db *postgres) {
	query := `create table if not exists cards (
		id serial primary key,
		number varchar(50) not null,
		userid integer not null,
		bankid integer not null,
		foreign key (userid) references users(id) on delete cascade on update cascade,
		foreign key (bankid) references banks(id) on delete cascade on update cascade,
		created_at timestamp
	)`

	_, err := db.Conn.Exec(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
}
