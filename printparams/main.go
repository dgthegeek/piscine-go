package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for _, arg := range args[1:] {
		for i, r := range arg {
			if i >= 0 {
				z01.PrintRune(r)
			}
		}
		z01.PrintRune('\n')
	}
}
