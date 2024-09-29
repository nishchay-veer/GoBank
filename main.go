package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/nishchay-veer/simplebank/api"
	db "github.com/nishchay-veer/simplebank/db/sqlc"
	"github.com/nishchay-veer/simplebank/util"
)


func main() {
	config,err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server , err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}

}
