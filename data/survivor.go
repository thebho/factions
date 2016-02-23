package data

import (
	"fmt"
	"log"
)

func InitializeSurvivor(db *DB) {
	fmt.Println("Initializing Survivor Table")
	// Test DB connection
	err := db.Ping()
	if err != nil {
		fmt.Printf("Ping DB error: %v\n", err)
		return
	}

	// Dynamically create table
	_, err = db.Exec("CREATE TABLE survivor ( uid integer, firstName varchar(32), lastName varchar(32) )")
	if err != nil {
		fmt.Printf("Skipping Create Table: %v\n", err)
		return
	}
	fmt.Println("Survivor Table Added")
}

func DeleteTable(db *DB) {
	fmt.Println("Deleting Survivor Table")
	_, err := db.Exec("DROP TABLE survivor")
	if err != nil {
		log.Printf(err.Error())
	}
	fmt.Println("Survivor Table Deleted")
}
