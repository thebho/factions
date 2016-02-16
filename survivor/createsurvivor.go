package survivor

import (
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

	fmt.Printf("Returning New Random Survivor: %v", randomSurvivor)
	return randomSurvivor
}

// TODO: Randomize name/sex generator
func generateRandomNameAndSex() (string, string, string) {
	fmt.Println("Generating Random Name/Sex")
	// Will return random generated name
	return "Random First", "Random Last", "(fe)male"
}

//TODO: Randomize Alignment
func generateRandomAlignment() (string, string) {
	fmt.Println("Generating Random Alignment Stats")
	return "Neutral", "Neutral"
}

//TODO: Randomize SurvivorStats
func generateRandomSurvivorStats() (int, int, int, int, int) {
	fmt.Println("Generating Random Survivor Stats")
	return 1, 2, 3, 4, 5
}
