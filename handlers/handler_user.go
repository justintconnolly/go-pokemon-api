package handlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

func generateAPIKey() (string, error) {
	bytes := make([]byte, 32) // Adjust key length as needed
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func storeAPIKey(db *sql.DB, key string) error {
	stmt, err := db.Prepare("INSERT INTO api_keys (key) VALUES ($1)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(key)
	return err
}

func GetAPIKey(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Generate a new API key
	newKey, err := generateAPIKey()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Store the key in the database
	if err := storeAPIKey(db, newKey); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Securely transmit the key to the user (e.g., using HTTPS and JSON response)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"key": newKey})
}
