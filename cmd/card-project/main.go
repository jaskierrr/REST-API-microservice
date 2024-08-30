package main

import (
	"card-project/internal/database"
	"card-project/internal/transport"
)

func main() {
	database.OpenConn()

	transport.Run()
}
