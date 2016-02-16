package survivor

import (
	"github.com/pborman/uuid"
)

// Survivor
type Survivor struct {
	Key             string
	Type            string
	FirstName       string
	LastName        string
	Sex             string
	Faction         string // Faction key
	Role            string // or job Hunter, Farmer, Politician, Teacher, Doctor, Defense
	FactionApproval int
	survivorStats
}

type Player struct {
	Survivor
	Experience int
}

type AISurvivor struct {
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
func NewSurvivor(firstName, lastName, sex string) *Survivor {
	return &Survivor{
		Key:       uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Sex:       sex,
		Type:      "Survivor",
	}
}

// NewAISurvivor
func NewAISurvivor(firstName, lastName, sex string) *AISurvivor {
	aisurvivor := &AISurvivor{}
	aisurvivor.Key = uuid.New()
	aisurvivor.FirstName = firstName
	aisurvivor.LastName = lastName
	aisurvivor.Sex = sex
	aisurvivor.Type = "AI Survivor"
	return aisurvivor
}

func (a *AISurvivor) AddAlignment(align, law string) {
	a.alignment = align
	a.lawalignment = law
}

func (s *Survivor) AddSurvivorStats(str, intel, leader, farm, hunt int) {
	s.strength = str
	s.intelligence = intel
	s.leadership = leader
	s.farming = farm
	s.hunting = hunt
}
