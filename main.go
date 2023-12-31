package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func connectToDatabase(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		fmt.Println("Could not connect to the database.")
		return nil, err
	}
	fmt.Println("Successfully connected to database!")
	return db, nil
}

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	dbConn, err := connectToDatabase(dbURL)
	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}
	defer dbConn.Close() // Ensure the connection is closed

	//api logic will go here and utilize the database connection called db
	http.HandleFunc("/api/v1/pokemon/", getPokemonByName) //TODO func from handler_pokemon.go undefined

	http.ListenAndServe(":8080", nil)
}
