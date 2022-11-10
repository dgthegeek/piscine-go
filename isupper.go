package piscine

func IsUpper(str string) bool {
	slice := []rune(str)

	for i := 0; i <= len(slice)-1; i++ {
		if slice[i] >= 0 && slice[i] <= 64 || slice[i] >= 91 && slice[i] <= 127 {
			return false
		}
	}
	return true
}
