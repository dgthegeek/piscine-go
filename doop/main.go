package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	nb1 := Atoi(args[0])
	signe := args[1]
	nb2 := Atoi(args[2])

	if len(args) != 3 || signe != "+" && signe != "*" && signe != "-" && signe != "/" && signe != "%" {
		// it prints nothing
	} else {
		if args[1] == "%" && nb2 == 0 {
			fmt.Print("No Modulo by 0\n")
		} else if args[1] == "/" && nb2 == 0 {
			fmt.Print("No division by 0\n")
		} else if args[1] == "+" {
			fmt.Println(nb1 + nb2)
		} else if args[1] == "-" {
			fmt.Println(nb1 - nb2)
		} else if args[1] == "*" {
			fmt.Println(nb1 * nb2)
		} else if args[1] == "/" {
			fmt.Println(nb1 / nb2)
		} else {
			fmt.Println(nb1 % nb2)
		}
	}
}
