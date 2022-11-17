package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := range a {
		if i+1 < len(a) {
			if f(a[i], a[i+1]) > 0 {
				return true
			} else if f(a[i], a[i+1]) < 0 {
				return true
			}
		}
	}
	return false
}
