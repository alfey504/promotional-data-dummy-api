package utils

import (
	"math/rand/v2"
	"time"
)

func RandomDate(start, end time.Time) time.Time {
	if start.After(end) {
		start, end = end, start // Swap if ordered incorrectly
	}

	// Calculate the difference in seconds between the two dates
	diff := end.Unix() - start.Unix()
	if diff <= 0 {
		return start
	}

	// Generate a random number of seconds within that difference
	randomSeconds := rand.Int64N(diff)

	// Add the random seconds back to the start date
	return start.Add(time.Duration(randomSeconds) * time.Second)
}
