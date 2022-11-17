package piscine

func IsSorted(f func(int, int) int, arr []int) bool {
	up := true
	down := true

	for i := 1; i < len(arr); i++ {
		if !(f(arr[i-1], arr[i]) >= 0) {
			up = false
		}
	}

	for i := 1; i < len(arr); i++ {
		if !(f(arr[i-1], arr[i]) <= 0) {
			down = false
		}
	}
	return up || down
}
