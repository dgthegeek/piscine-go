package main

import (
	"github.com/01-edu/z01"
)

func upper(a rune) rune {
	return (((a - 'A') + 1) % 26) + 'A'
}

func lower(a rune) rune {
	return (((a - 'a') + 1) % 26) + 'a'
}

func Rot14(s string) string {
	var result string

	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			r = upper(r)
		} else if r >= 'a' && r <= 'z' {
			r = lower(r)
		}

		result = result + string(r)
	}
	return result
}

func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
