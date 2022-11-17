package piscine

func Abort(a, b, c, d, e int) int {
	arg := []int{a, b, c, d, e}
	SortIntegerTable(arg)
	return arg[2] // Will get the 3rd number
}
