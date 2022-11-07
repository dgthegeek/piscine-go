package piscine

func BasicAtoi(s string) int {
	slice := []rune(s)
	n := len(s)
	var transformed int
	for i := 0; i < n; i++ {
		if slice[i] < '0' || slice[i] > '9' {
			return 0
		} else {
			transformed += int(slice[i]) - '0'
		}
	}
	return transformed
}
