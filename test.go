package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if !(args[0] >= "0" && args[2] <= "9") || !(args[0] >= "0" && args[2] <= "9") {
		return
	}
	value1, _ := strconv.Atoi(args[0])
	value2, _ := strconv.Atoi(args[2])
	signe := args[1]
	if signe != "+" && signe != "-" && signe != "/" && signe != "%" {
		fmt.Println("signe invalide")
	}
	switch signe {
	case "+":
		os.Stdout.Write([]byte(value1 + value2))

	case "-":
		os.Stdout.Write([]byte(value1 - value2))
	case "/":
		if value2 == 0 {
			os.Stdout.Write([]byte("0 Division error"))
			os.Exit((1))
		}
		os.Stdout.Write([]byte(value1 / value2))
	case "%":
		if value2 == 0 {
			os.Stdout.Write([]byte("Modulo 0 error"))
			os.Exit((1))
		}
		os.Stdout.Write([]byte(value1 - value2))
	case "*":
		os.Stdout.Write([]byte(value1 * value2))

	}
}
