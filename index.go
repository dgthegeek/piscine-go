package piscine

func Index(s string, toFind string) int {
	a := []rune(s)
	b := []rune(toFind)
	len_a := len(a)
	len_b := len(b)

	for i := 0; i <= len_a-len_b; i++ { // because the lenght of the Tofind word will never be greater than len_a - len_b
		if toFind == s[i:i+len_b] {
			return (i)
		}
	}
	return -1
}
