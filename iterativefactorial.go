package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	limit := 21
	factorial := 1
	for i := 1; i <= nb; i++ {
		if i > limit {
			return 0
		}
		factorial *= i
	}
	return factorial
}
