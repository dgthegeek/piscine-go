package piscine

func ToLower(s string) string {
	a := []rune(s)
	result := ""

	for i, value := range a {
		if value >= 'a' && value <= 'z' {
			value = value
		} else if value >= 'A' && value <= 'Z' {
			a[i] = a[i] + 32
		}
		result += string(a[i])
	}
	return result
}
