package types

import (
	"github.com/pborman/uuid"
)

type Survivor struct {
	Key       string
	Type      string
	FirstName string
	LastName  string
	Faction   string // Faction key
	SurvivorStats
}

type Player struct {
	Survivor
}

type SurvivorStats struct {
	Strength     int
	Intelligence int
	Leadership   int
	Farming      int
	Hunting      int
}

func NewSurvivor(firstName, lastName string) *Survivor {
	return &Survivor{
		Key:       uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Type:      "Survivor",
	}
}
