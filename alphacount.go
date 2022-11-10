package piscine

func AlphaCount(s string) int {
	slice := []rune(s)
	n := len(slice)
	counter := 0
	for i := 0; i < n; i++ {
		if (slice[i] >= 'a' && slice[i] <= 'z') || (slice[i] >= 'A' && slice[i] <= 'Z') {
			counter++
		}
	}
	return counter
}
