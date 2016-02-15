package main

import (
	"factions/types"
	"fmt"
)

func main() {
	survivor := types.NewSurvivor("Brian", "Hoehne")
	faction := types.NewFaction("Test Faction")
	faction.Leader = "LeaderKey"
	faction.Members = []string{"LeaderKey", "Other Member"}
	garden := types.NewVegetableGarden("Test Garden")
	faction.FoodSources = append(faction.FoodSources, garden)
	garden.AddSurvivor(*survivor)
	garden.GatherFood(faction)
	fmt.Println(faction.TotalFood)
	fmt.Printf("Faction: %+v\n", faction.Members)
	fmt.Printf("Garden: %+v\n", garden)
	fmt.Printf("Survivor: %+v\n", survivor)

}
