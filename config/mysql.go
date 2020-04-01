package config

import (
	"database/sql"
	"log"
)

// Connect export functions
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang1")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

// Testing export functions
func Testing() string { // Karena function ini akan di ekspor, maka penulisan harus kapital
	return "\nThis from another file"
}
