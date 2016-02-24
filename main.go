package main

import (
	"factions/data"
	"factions/faction"
	"factions/random"
	"factions/survivor"
	"fmt"
)

func main() {
	db, err := data.ConnectToDatabase()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Database")
	defer db.Close()
	// data.DeleteTable(db)
	data.DeleteTable(db)

	data.InitializeSurvivor(db)
	testFactions(db)
	data.DeleteAllRows(db)
	data.DeleteTable(db)
	// testRandom()

}

func testFactions(db *data.DB) {
	for i := 0; i < 10; i++ {
		newSurvivor := survivor.CreateRandomSurivor()
		data.AddNewSurvivor(db, newSurvivor)
		newFaction := faction.NewFaction("Test Faction")
		newFaction.Leader = "LeaderKey"
		newFaction.Members = []string{"LeaderKey", "Other Member"}
		garden := faction.NewVegetableGarden("Test Garden")
		newFaction.FoodSources = append(newFaction.FoodSources, garden)
		garden.AddSurvivor(newSurvivor.Key)
		garden.GatherFood(newFaction)
		// fmt.Println(newFaction.TotalFood)
		// fmt.Printf("Faction: %+v\n", newFaction.Members)
		// fmt.Printf("Garden: %+v\n", garden)
		fmt.Printf("Survivor: %+v\n", newSurvivor)
		// 	fmt.Printf("Survivor: %s, STR: %d INT: %d\n", newSurvivor.FirstName, newSurvivor.GetStat(survivor.SUR_STR), newSurvivor.GetStat(survivor.SUR_INT))
		// 	newSurvivor.AdjustStat(survivor.SUR_STR, 3)
		// 	newSurvivor.AdjustStat(survivor.SUR_INT, -3)
		// 	fmt.Printf("Survivor: %s, STR: %d INT: %d\n", newSurvivor.FirstName, newSurvivor.GetStat(survivor.SUR_STR), newSurvivor.GetStat(survivor.SUR_INT))
		// }
	}
}

func testRandom() {
	for i := 0; i < 100; i++ {
		fmt.Println(random.RandInt(10))
	}
}
