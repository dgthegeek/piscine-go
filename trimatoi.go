package piscine

func TrimAtoi(s string) int {
	runes := []rune(s)
	sign := 1
	resultat := 0
	foundDigit := false
	for _, value := range runes {
		if !foundDigit {
			if value == '-' {
				sign = -1
			} else if value == '+' {
				sign = 1
			}
		}
		if value >= '0' && value <= '9' {
			foundDigit = true // the signe is just next to this digits
			preresultat := 10 * resultat
			resultat = preresultat + int(value) - '0'
		}
	}
	return sign * resultat
}
