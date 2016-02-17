package faction

import (
	"fmt"
)

// Implements FoodSource
type VegetableGarden struct {
	FoodSourceStats
}

func NewVegetableGarden(name string) *VegetableGarden {
	garden := &VegetableGarden{}
	garden.Type = "Vegetable Garden"
	garden.HarvestPerDay = 3
	garden.Name = name
	return garden
}

func (v *VegetableGarden) AddSurvivor(survivor string) {
	v.Survivors = append(v.Survivors, survivor)
}

func (v *VegetableGarden) GatherFood(faction *Faction) {
	foodGathered := v.CalculateFoodGathered()
	faction.TotalFood += foodGathered
	fmt.Printf("Harvesting %d food.  Faction total food: %d \n", foodGathered, faction.TotalFood)
}
