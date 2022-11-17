package piscine

func IsSorted(f func(int, int) int, a []int) bool {
	down := false
	up := false

	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) >= 0 {
			down = true
		}
	}

	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) <= 0 {
			up = true
		}
	}
	return down || up
}
