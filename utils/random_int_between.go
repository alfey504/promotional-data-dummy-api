package utils

import "math/rand/v2"

func RandomIntBetween(x int, y int) int {
	// If x is greater than y, swap them to prevent errors
	if x > y {
		x, y = y, x
	}

	// Returns a number in the range [x, y]
	return rand.IntN(y-x+1) + x
}
