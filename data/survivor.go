package data

import (
	"factions/survivor"
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
	_, err = db.Exec("CREATE TABLE survivor ( surivorkey varchar(40), firstName varchar(32), lastName varchar(32), sex varchar(10), faction varchar(36), role varchar(36), factionApproval integer )")
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

func DeleteAllRows(db *DB) {
	fmt.Println("Deleting Survivor Table Rows")
	_, err := db.Exec("DELETE FROM survivor")
	if err != nil {
		log.Printf(err.Error())
	}

}

func AddNewSurvivor(db *DB, survivor *survivor.AISurvivor) {
	// Modifying Data
	fmt.Println("Adding New Survivor")
	stmt, err := db.Prepare("INSERT INTO survivor(surivorkey, firstName, lastName, sex, faction, role, factionApproval) VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Printf(err.Error())
	}

	fmt.Println("Inserting Data")
	fmt.Printf("Survivor: %+v\n", survivor)
	res, err := stmt.Exec(survivor.Key, survivor.FirstName, survivor.LastName, survivor.Sex, survivor.Faction, survivor.Role, survivor.FactionApproval)
	if err != nil {
		log.Printf(err.Error())
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Printf("ID: %d\n", id)

	fmt.Println("New Survivor Added")
}

func QuerySurvivor(db *DB) (*survivor.AISurvivor, error) {
	fmt.Println("Querying Data")
	rows, err := db.Query("select * from survivor")
	if err != nil {
		fmt.Printf("Query Err: %v\n", err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			survivorKey string
			firstName   string
			lastName    string
			faction     string
			role        string
			sex         string
		)
		err := rows.Scan(&survivorKey, &firstName, &lastName, &role, &faction, &sex)
		if err != nil {
			fmt.Printf("Scan Err: %v\n", err)
			return nil, err
		}
		fmt.Printf("Name: %s Key: %s Sex: %s\n", firstName+lastName, survivorKey, sex)
		fmt.Printf("Faction: %s Role: %s\n", faction, role)

		survivor := &survivor.AISurvivor{}
		survivor.Key = survivorKey
		survivor.Type = "Survivor"
		survivor.FirstName = firstName
		survivor.LastName = lastName
		survivor.Sex = sex
		survivor.Faction = faction
		survivor.Role = role
		return survivor, nil
	}
	err = rows.Err()
	if err != nil {
		log.Printf("Rows err %s\n", err)
		return nil, err
	}
	return nil, fmt.Errorf("Unknown err")
}
