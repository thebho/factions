package faction

type FoodSource interface {
	GatherFood(*Faction) //Number of survivors
	AddSurvivor(string)
}

type FoodSourceStats struct {
	Type          string
	Name          string
	HarvestPerDay int
	Survivors     []string
}
