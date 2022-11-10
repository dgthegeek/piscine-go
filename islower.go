package piscine

func IsLower(str string) bool {
	slice := []rune(str)

	for i := 0; i <= len(slice)-1; i++ {
		if slice[i] >= 0 && slice[i] <= 96 || slice[i] >= 123 && slice[i] <= 127 {
			return false
		}
	}
	return true
}
