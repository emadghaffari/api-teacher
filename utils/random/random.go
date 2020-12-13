package random

import (
	"math/rand"
	"time"
)

// Slice random Slice of numbers
func Slice(num, min, max int) []int {
	if min > max {
		return nil
	}
	rand.Seed(time.Now().UnixNano())
	var items []int
	for i := 0; i < num; i++ {
		random := (rand.Intn((max - min)) + min)
		items = append(items, random)
	}
	return items
}

// Rand between numbers
func Rand(min, max int) int {
	if min > max {
		return 0
	}
	rand.Seed(time.Now().UnixNano())
	random := (rand.Intn((max - min)) + min)
	return random
}
