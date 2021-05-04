package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/thanhftu/simple_bank/api"
	db "github.com/thanhftu/simple_bank/db/sqlc"
	"github.com/thanhftu/simple_bank/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to DB", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddr)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
