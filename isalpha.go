package piscine

func IsAlpha(str string) bool {
	for _, value := range str {
		if value >= 'a' && value <= 'z' || value >= 'A' && value <= 'Z' || value >= '0' && value <= '9' {
			return true
		}
	}
	return false
}
