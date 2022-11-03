package main

import "github.com/01-edu/z01"

func main() {
	var aRune string = "abcdefghijklmnopqrstuvwxyz"

	for i := range [26]int{} {
		z01.PrintRune(rune(aRune[i]))
	}
	z01.PrintRune('\n')
}
