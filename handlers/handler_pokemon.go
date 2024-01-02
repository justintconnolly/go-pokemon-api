package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"unicode"

	"github.com/justintconnolly/pokemon-api/models"
)

func GetPokemonByName(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed on this endpoint", http.StatusMethodNotAllowed)
		return
	}
	// Get the Pokemon name from the request
	requestedName := r.URL.Query().Get("name")
	if requestedName == "" {
		http.Error(w, "Name parameter is required", http.StatusBadRequest)
		return
	}
	runes := []rune(requestedName)
	if len(runes) > 0 {
		runes[0] = unicode.ToUpper(runes[0])
	}
	pokemonName := string(runes)

	// Execute the SQL query to fetch the Pokemon data
	row := db.QueryRow("SELECT name, pokedex_number, type1, generation FROM pokemon WHERE name = $1", pokemonName)

	// Scan the row into a Pokemon struct
	var pokemon models.PokemonDefaultResponse
	err := row.Scan(&pokemon.Name, &pokemon.PokedexNumber, &pokemon.Type1, &pokemon.Generation) // Scan into respective fields

	// Handle query errors
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
