package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/brianvoe/gofakeit/v7"
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
		db, err := sql.Open("postgres", dbUrl)
		if err != nil {
				log.Fatal(err)
		}
		defer db.Close()
 
		 // Test the connection
		 err = db.Ping()
		 if err != nil {
				 log.Fatal(err)
		 }
 
		 fmt.Println("Successfully connected to PostgreSQL database!")
 
		 // Example of a simple query
		 rows, err := db.Query("SELECT id, name FROM users LIMIT 5")
		 if err != nil {
				 log.Fatal(err)
		 }
		 defer rows.Close()
 
		 // Iterate through results
		 for rows.Next() {
				 var id int
				 var name string
				 if err := rows.Scan(&id, &name); err != nil {
						 log.Fatal(err)
				 }
				 fmt.Printf("ID: %d, Name: %s\n", id, name)
		 }
 
		 // Check for any errors encountered during iteration
		 if err = rows.Err(); err != nil {
				 log.Fatal(err)
		 }

	fmt.Println(gofakeit.Name())
	fmt.Println(gofakeit.LastName())
	fmt.Println(gofakeit.FirstName())
	fmt.Println(gofakeit.PastDate().UTC())
	fmt.Println(gofakeit.FutureDate())
	fmt.Println(gofakeit.BuzzWord())
	fmt.Println(gofakeit.Company())
	fmt.Println(gofakeit.HackerPhrase())
	fmt.Println(gofakeit.HackerNoun())
	fmt.Println(gofakeit.BS())
	fmt.Println(gofakeit.BeerStyle())
}