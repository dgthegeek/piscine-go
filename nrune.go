package piscine

func NRune(s string, n int) rune {
	slice := []rune(s)
	for i, r := range slice {
		if i+1 == n {
			return r
		}
	}
	return 0
}
