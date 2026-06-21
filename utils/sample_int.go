package utils

func SampleRandInt(x, y, n int) []int {
	slice := make([]int, n)

	if x > y {
		x, y = y, x
	}
	for i := 0; i < n; i++ {
		slice[i] = RandomIntBetween(x, y)
	}

	return slice
}
