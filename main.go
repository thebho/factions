package main

import (
	"factions/faction"
	"factions/survivor"
	"fmt"
)

func main() {
	survivor := survivor.CreateRandomSurivor()
	newFaction := faction.NewFaction("Test Faction")
	newFaction.Leader = "LeaderKey"
	newFaction.Members = []string{"LeaderKey", "Other Member"}
	garden := faction.NewVegetableGarden("Test Garden")
	newFaction.FoodSources = append(newFaction.FoodSources, garden)
	garden.AddSurvivor(survivor.Key)
	garden.GatherFood(newFaction)
	fmt.Println(newFaction.TotalFood)
	fmt.Printf("Faction: %+v\n", newFaction.Members)
	fmt.Printf("Garden: %+v\n", garden)
	fmt.Printf("Survivor: %+v\n", survivor)

}
