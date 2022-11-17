package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	nb1 := Atoi(args[0])
	signe := args[1]
	nb2 := Atoi(args[2])

	// errors
	mod := "No Modulo by 0"
	Lmod := []rune(mod)
	div := "No division by 0"
	Ldiv := []rune(div)

	if len(args) != 3 || signe != "+" && signe != "*" && signe != "-" && signe != "/" && signe != "%" {
		// it prints nothing
	} else {
		if args[1] == "%" && nb2 == 0 {
			for _, v := range Lmod {
				z01.PrintRune(v)
			}
		} else if args[1] == "/" && nb2 == 0 {
			for _, v := range Ldiv {
				z01.PrintRune(v)
			}
		} else if args[1] == "+" {
			z01.PrintRune((nb1 + nb2) + '0')
		} else if args[1] == "-" {
			z01.PrintRune((nb1 - nb2) + '0')
		} else if args[1] == "*" {
			z01.PrintRune((nb1 * nb2) + '0')
		} else if args[1] == "/" {
			z01.PrintRune((nb1 / nb2) + '0')
		} else {
			z01.PrintRune((nb1 % nb2) + '0')
		}
	}
}
