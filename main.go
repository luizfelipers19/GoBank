package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/luizfelipers19/gobank/api"
	db "github.com/luizfelipers19/gobank/db/sqlc"
	"github.com/luizfelipers19/gobank/util"
)

// const (
// 	dbDriver      = "postgres"
// 	dbSource      = "postgresql://root:secret@localhost:5432/gobank?sslmode=disable"
// 	serverAddress = "localhost:8080"
// )

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can not load stored configs ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can not connect to the Database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Can not start server:", err)
	}
}
