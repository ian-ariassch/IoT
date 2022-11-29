package db

import (
	"context"
	"log"
	"minimal/ent"

	_ "github.com/lib/pq"
)

var postgresClient *ent.Client

func NewPostgresClient() *ent.Client {
	if postgresClient != nil {
		return postgresClient
	}

	client, err := ent.Open("postgres", "user=postgres password=[HERE_YOUR_PASSWORD] host=[HERE_YOUR_HOST] port=5432 dbname=postgres")
	if err != nil {
		log.Printf("failed opening connection to postgres: %v", err)
		panic(err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Println(err)
		panic(err)
	}

	postgresClient = client

	return postgresClient
}
