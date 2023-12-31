package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/justintconnolly/pokemon-api/models"
)

func getPokemonByName(w http.ResponseWriter, r *http.Request) {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}
	// Get the Pokemon name from the request
	pokemonName := r.URL.Query().Get("name")

	// Connect to the database
	db, err := connectToDatabase(dbURL)
	if err != nil {
		fmt.Println("Could not connect to the database.")
		return
	}

	// Execute the SQL query to fetch the Pokemon data
	row := db.QueryRow("SELECT name, pokedex_number, type1, generation FROM pokemon WHERE name = $1", pokemonName)

	// Scan the row into a Pokemon struct
	var pokemon models.Pokemon
	row.Scan(&pokemon.Name, &pokemon.PokedexNumber, &pokemon.Type1, &pokemon.Generation) // Scan into respective fields

	// Handle errors
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Pokemon not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	// Serialize the Pokemon data into JSON
	jsonData, err := json.Marshal(pokemon)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
