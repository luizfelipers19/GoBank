package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/luizfelipers19/gobank/api"
	db "github.com/luizfelipers19/gobank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/gobank?sslmode=disable"
	serverAddress = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can not connect to the Database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Can not start server:", err)
	}
}
