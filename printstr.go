package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	my_slice := []rune(s)
	for i := 0; i < len(my_slice); i++ {
		z01.PrintRune(my_slice[i])
	}
}
