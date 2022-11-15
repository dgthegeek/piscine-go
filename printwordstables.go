package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, mot := range a {
		for _, letter := range mot {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
