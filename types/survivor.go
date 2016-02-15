package types

import (
	"github.com/pborman/uuid"
)

// Survivor
type Survivor struct {
	Key             string
	Type            string
	FirstName       string
	LastName        string
	Faction         string // Faction key
	Role            string // or job Hunter, Farmer, Politician, Teacher, Doctor, Defense
	FactionApproval int
	survivorStats
}

type Player struct {
	Survivor
	Experience int
}

type AISurivor struct {
	Survivor
	aiStats
}

type survivorStats struct {
	strength     int
	intelligence int
	leadership   int
	farming      int
	hunting      int
}

type aiStats struct {
	alignment    string //Good, Neutral, Evil
	lawalignment string //Lawful, Neutral, Chaotic

}

// NewSurvivor
func NewSurvivor(firstName, lastName string) *Survivor {
	return &Survivor{
		Key:       uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Type:      "Survivor",
	}
}
