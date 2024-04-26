package main

import (
	"context"
	"log"

	"github.com/Dezmond-sama/Specialist_Go_Courses/GO3/bankstore/api"
	db "github.com/Dezmond-sama/Specialist_Go_Courses/GO3/bankstore/db/sqlc"
	"github.com/Dezmond-sama/Specialist_Go_Courses/GO3/bankstore/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	pool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	store := db.NewStore(pool)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal(err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal(err)
	}
}
