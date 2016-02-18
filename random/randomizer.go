package random

import (
	"math/rand"
)

// Returns a number betweek 1 and the high argument
func RandInt(high int) int {
	return 1 + rand.Intn(high)
}

// Returns an index between 0 and the count (non-inclusive)
func RandIndex(count int) int {
	return rand.Intn(count - 1)
}
