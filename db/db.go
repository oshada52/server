package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Global db variable

var Connection *sql.DB

// Create a new instance of DB

func InitDB() {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		"postgres", "Dwoshada", "pg-db", 5432, "postgres")

	var err error
	Connection, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

}
