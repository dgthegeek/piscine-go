package piscine

func LastRune(s string) rune {
	slice := []rune(s)
	lenght := len(slice)
	return (slice[lenght-1])
}
