package piscine

func Map(f func(int) bool, a []int) []bool {
	var slice []bool
	for _, v := range a {
		slice = append(slice, f(v))
	}
	return slice
}
