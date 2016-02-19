package random

import (
	"math/rand"
	"time"
)

// Returns a number betweek 1 and the high argument
func RandInt(high int) int {
	return randInt(1, high)
}

// Returns an index between 0 and the count (non-inclusive)
func RandIndex(count int) int {
	return randInt(0, count)
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
