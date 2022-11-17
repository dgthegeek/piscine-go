package piscine

func IsSorted(f func(int, int) int, a []int) bool {
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) >= 0 {
			return true
		}
	}

	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) <= 0 {
			return true
		}
	}
	return false
}
