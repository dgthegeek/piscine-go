package piscine

import "github.com/01-edu/z01"

func StrRev(s string) string {
	my_slice := []rune(s)
	for i := len(my_slice); i == 0; i-- {
		z01.PrintRune(my_slice[i])
	}
	return ""
}
