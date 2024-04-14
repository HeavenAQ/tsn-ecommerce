package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	// load .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Panicf("Error loading .env file: %v\n", err)
	}

	ctx := context.Background()
	dbSource := os.Getenv("DB_SOURCE")
	testDBPool, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Printf("Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer testDBPool.Close()
	testQueries = New(testDBPool)
	m.Run()
}
