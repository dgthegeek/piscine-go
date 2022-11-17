package piscine

func SortIntegerTable(table []int) {
	tampon := 0
	for i := 0; i < len(table)-1; i++ {
		for j := 0; j < len(table)-1; j++ {
			if table[j] > table[j+1] {
				tampon = table[j]
				table[j] = table[j+1]
				table[j+1] = tampon

			}
		}
	}
}

func Abort(a, b, c, d, e int) int {
	arg := []int{a, b, c, d, e}
	SortIntegerTable(arg)
	return arg[2] // Will get the 3rd number
}
