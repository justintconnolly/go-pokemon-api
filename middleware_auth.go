package main

import (
	"database/sql"
	"net/http"
)

func isValidAPIKey(db *sql.DB, key string) bool {
	var exists bool
	err := db.QueryRow("SELECT EXISTS( SELECT 1 FROM api_keys WHERE key = $1 AND expires_at IS NULL )", key).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func validateAPIKey(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key") // Assuming the API key is in this header

		// Check if the key exists and is valid
		if apiKey == "" || !isValidAPIKey(db, apiKey) { // Replace with your validation logic
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		} // Key is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
