package types

import "fmt"

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

func (v *VegetableGarden) AddSurvivor(survivor Survivor) {
	fmt.Printf("Adding Survivor: %s\n", survivor.FirstName)
	v.Survivors = append(v.Survivors, survivor.Key)
}

func (v *VegetableGarden) GatherFood(faction *Faction) {
	fmt.Printf("Harvesting %d food.  Faction total food: %d ", v.HarvestPerDay*len(v.Survivors), faction.TotalFood)
	faction.TotalFood += v.HarvestPerDay * len(v.Survivors)
}
