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

// CalculateFoodGathered uses the survivors keys to decide how much food is gathered based on surivors skills and HarvestPerDay
func (f *FoodSourceStats) CalculateFoodGathered() int {
	// TODO: Pull each survivor's Farming stats from DB and multiply by HarvestPerDay
	return -1
}
