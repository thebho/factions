package survivor

import (
	"fmt"

	rand "github.com/thebho/random-tools"
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

func generateRandomNameAndSex() (string, string, string) {
	fmt.Println("Generating Random Name/Sex")
	// Will return random generated name
	sexInt := rand.RandIndex(2)
	fmt.Println(sexInt)
	var first string
	if sexInt == 0 {
		first = male[rand.RandIndex(len(male))]
	} else {
		first = female[rand.RandIndex(len(male))]
	}
	last := surname[rand.RandIndex(len(surname))]
	return first, last, sex[sexInt]
}

func generateRandomAlignment() (string, string) {
	fmt.Println("Generating Random Alignment Stats")
	return alignment[rand.RandIndex(len(alignment))], lawalignment[rand.RandIndex(len(lawalignment))]
}

func generateRandomSurvivorStats() (int, int, int, int, int) {
	fmt.Println("Generating Random Survivor Stats")
	return rand.RandInt(99), rand.RandInt(99), rand.RandInt(99), rand.RandInt(99), rand.RandInt(99)
}
