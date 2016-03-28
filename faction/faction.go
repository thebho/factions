package faction

import (
	"github.com/pborman/uuid"
)

//Faction describes
type Faction struct {
	ID      string   `db:"uid"`
	Name    string   `db:"name"`
	Type    string   `db:"type"`
	Leader  string   `db:"leader"`
	Members []string //Survivor Keys
	FactionConsumeables
}

//NewFaction is a Faction contructor
func NewFaction(name string) *Faction {
	return &Faction{
		ID:   uuid.New(),
		Name: name,
		Type: "Faction",
	}
}

type FactionPolitics struct {
	PoliticalType string //Dictator/Democracy/Socialist
}

type FactionConsumeables struct {
	FoodSources []FoodSource
	TotalFood   int //That has been moved from FoodSources to FoodStored
}
