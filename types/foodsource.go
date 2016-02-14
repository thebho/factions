package types

type FoodSource interface {
	GatherFood(*Faction) //Number of survivors
	AddSurvivor(Survivor)
}

type FoodSourceStats struct {
	Type          string
	Name          string
	HarvestPerDay int
	Survivors     []string
}
