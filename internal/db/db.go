package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

var Pool *pgxpool.Pool

func InitDB(connString string) {
	var err error
	Pool, err = pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
}

func CloseDB() {
	Pool.Close()
}
