package piscine

func ToUpper(s string) string {
	a := []rune(s)
	result := ""

	for i, value := range a {
		if value >= 'A' && value <= 'Z' {
			value = value
		} else if value >= 'a' && value <= 'z' {
			a[i] = a[i] - 32
		}
		result += string(a[i])
	}
	return result
}
