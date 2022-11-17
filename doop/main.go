package main

import (
	"os"
)

func Atoi(s string) int {
	tab := []rune(s)
	a := 0
	if len(tab) == 0 {
		return 0
	}
	for i := 0; i < len(tab); i++ {
		if tab[i] == 45 || tab[i] == 43 || (tab[i] >= 48 && tab[i] <= 57) {
			if i != 0 && (tab[i] == 45 || tab[i] == 43) {
				return 0
			}
			if tab[i] >= 48 && tab[i] <= 57 {
				a *= 10
				a = a + int(tab[i]) - '0'
			}

		} else {
			return 0
		}
	}
	if tab[0] == 45 {
		a *= -1
	}

	return a
}

func ff(args []string) string {
	nb1 := Atoi(args[0])
	signe := args[1]
	nb2 := Atoi(args[2])

	// errors

	if len(args) != 3 || signe != "+" && signe != "*" && signe != "-" && signe != "/" && signe != "%" {
		// it prints nothing
	} else {
		if args[1] == "%" && nb2 == 0 {
			return "No Modulo by 0"
		} else if args[1] == "/" && nb2 == 0 {
			return "No division by 0"
		} else if args[1] == "+" {
			return PrintNbr(nb1 + nb2)
		} else if args[1] == "-" {
			return PrintNbr(nb1 - nb2)
		} else if args[1] == "*" {
			return PrintNbr(nb1 * nb2)
		} else if args[1] == "/" {
			return PrintNbr(nb1 / nb2)
		} else {
			return PrintNbr(nb1 % nb2)
		}
	}
	return "0"
}

func main() {
	args := os.Args[1:]
	ff(args)
}
