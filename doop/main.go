package main

import (
	"os"
)

func ff() string {
	args := os.Args[1:]
	nb1 := atoi(args[0])
	signe := args[1]
	nb2 := atoi(args[2])

	// errors

	if len(args) != 3 || signe != "+" && signe != "*" && signe != "-" && signe != "/" && signe != "%" {
		// it prints nothing
	} else {
		if args[1] == "%" && nb2 == 0 {
			return "No Modulo by 0"
		} else if args[1] == "/" && nb2 == 0 {
			return "No division by 0"
		} else if args[1] == "+" {
			return ((nb1 + nb2) + '0')
		} else if args[1] == "-" {
			return ((nb1 - nb2) + '0')
		} else if args[1] == "*" {
			return string((nb1 * nb2) + '0')
		} else if args[1] == "/" {
			return ((nb1 / nb2) + '0')
		} else {
			return ((nb1 % nb2) + '0')
		}
	}
	return "0"
}

func main() {
	ff()
}
