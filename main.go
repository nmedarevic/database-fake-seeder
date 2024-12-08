package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")

		return
	}

	// Retrieve environment variable with existence check
	dbUrl, exists := os.LookupEnv("DB_URL")
	if !exists {
		log.Fatal("DB_URL is not set")
	}

	// Open a connection to the database
	db, err := pgx.Connect(context.Background(), dbUrl)

	// db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	// Test the connection
	err = db.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to PostgreSQL database!")

	rowsData := [][]interface {
	}{
		{uuid.New().String(), gofakeit.Email(), gofakeit.Password(false, false, false, false, false, 0), gofakeit.PastDate(), gofakeit.PastDate(), gofakeit.Username(), gofakeit.Name()},
		{uuid.New().String(), gofakeit.Email(), gofakeit.Password(false, false, false, false, false, 0), gofakeit.PastDate(), gofakeit.PastDate(), gofakeit.Username(), gofakeit.Name()},
	}

	numberOfRowsInserted, err := db.CopyFrom(
		context.Background(),
		pgx.Identifier{"user"},
		[]string{"id", "email", "password", "created_at", "updated_at", "username", "full_name"},
		pgx.CopyFromRows(rowsData),
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Rows inserted %d\n", numberOfRowsInserted)
}
