package types

import (
	"fmt"
	"github.com/pborman/uuid"
)

type Faction struct {
	Key     string
	Name    string
	Type    string
	Leader  string   // Survivor key
	Members []string //Survivor Keys
	FactionConsumeables
}

func NewFaction(name string) *Faction {
	return &Faction{
		Key:  uuid.New(),
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
