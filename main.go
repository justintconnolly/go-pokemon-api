package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func connectToDatabase() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", user, password, dbname, sslmode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Could not connect to the database.")
		return nil, err
	}
	fmt.Println("Successfully connected to database!")
	return db, nil
}

func main() {
	fmt.Println("Best Pokemon API ever ;)")

	connStr := "user=your_username password=your_password dbname=your_database sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close() // Ensure the connection is closed

}
