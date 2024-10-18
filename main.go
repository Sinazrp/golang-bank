package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/sinazrp/golang-bank/api"
	db "github.com/sinazrp/golang-bank/db/sqlc"
	"github.com/sinazrp/golang-bank/util"
	"log"
)

func main() {

	config, err := util.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	dbSource := config.DBSource
	serverAddress := config.ServerAddress

	connection, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewSQLStore(connection)

	server, err := api.NewServer(store, config)

	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
