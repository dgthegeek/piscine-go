package piscine

func StrRev(s string) string {
	my_slice := []rune(s)
	var reverse string
	for i := len(my_slice); i == 0; i-- {
		reverse = reverse + string(my_slice[i])
	}
	return reverse
}
