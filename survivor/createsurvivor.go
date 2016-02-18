package survivor

import (
	rand "factions/random"
	"fmt"
)

func CreateRandomSurivor() *AISurvivor {
	fmt.Println("Creating a New Random Survivor")
	first, last, sex := generateRandomNameAndSex()
	randomSurvivor := NewAISurvivor(first, last, sex)

	align, law := generateRandomAlignment()
	randomSurvivor.AddAlignment(align, law)

	s, i, l, f, h := generateRandomSurvivorStats()
	randomSurvivor.AddSurvivorStats(s, i, l, f, h)

	fmt.Printf("Returning New Random Survivor: %v\n", randomSurvivor)
	return randomSurvivor
}

// TODO: Randomize names
func generateRandomNameAndSex() (string, string, string) {
	fmt.Println("Generating Random Name/Sex")
	// Will return random generated name
	return "Random First", "Random Last", sex[rand.RandIndex(2)]
}

func generateRandomAlignment() (string, string) {
	fmt.Println("Generating Random Alignment Stats")
	return alignment[rand.RandIndex(len(alignment))], lawalignment[rand.RandIndex(len(lawalignment))]
}

//TODO: Randomize SurvivorStats
func generateRandomSurvivorStats() (int, int, int, int, int) {
	fmt.Println("Generating Random Survivor Stats")
	return rand.RandInt(99), rand.RandInt(99), rand.RandInt(99), rand.RandInt(99), rand.RandInt(99)
}
