package piscine

func IsAlpha(s string) bool {
	slice := []rune(s)
	for _, value := range slice {
		if !(value >= 97 && value <= 122) && !(90 <= value && value >= 65) {
			return false
		}
	}
	return true
}
