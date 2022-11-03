package main

import "github.com/01-edu/z01"

func main() {
	var aRune string = "abcdefghijklmnopqrstuvwxyz"

	for i := len(aRune) - 1; i >= 0; i-- {
		z01.PrintRune(rune(aRune[i]))
	}
	z01.PrintRune('\n')
}
