package data

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	Survivor = "survivor"
)

// TODO: get the IP & Database Name from KV Store
var (
	tempServer   = "demo:demo@(127.0.0.1:3306)/"
	tempDatabase = "factionsTemp"
)

func ConnectToDatabase() (*sql.DB, error) {
	fmt.Println("Connecting to database")
	db, err := sql.Open("mysql", tempServer)
	if err != nil {
		return nil, err
	}

	// Create new database dynamically
	fmt.Println("Creating Database")
	_, err = db.Exec("CREATE DATABASE " + tempDatabase)
	if err != nil {
		fmt.Printf("Database exists\n")
	}

	// Select table to use
	_, err = db.Exec("USE " + tempDatabase)
	if err != nil {
		return nil, fmt.Errorf("Use: %v", err)
	}

	fmt.Println("Pinging database")
	// Test DB connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Ping: %v", err)
	}
	fmt.Println("Returning database")

	return db, nil
}
