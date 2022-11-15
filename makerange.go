package piscine

func MakeRange(min, max int) []int {
	var slice []int

	if min >= max {
		return slice
	} else {
		n := max - min
		slice = make([]int, n)
		for i := 0; i < n; i++ {
			slice[i] = min + i
		}
		return slice
	}
}
