package main

import (
	"fmt"
	"strings"

	"github.com/01-edu/z01"
)

func Reverse(s string) string {
	var reverse string
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}

func PrintNbrBase(nbr int, str string) {
	indx := 0
	for _, res := range str {
		if string(res) == "-" || string(res) == "+" || strings.Count(str, string(res)) > 1 {
			indx = 1
			break
		}
	}
	if indx == 1 || len(str) < 2 {
		fmt.Print("NV")
	} else if 2147483647 < nbr || -2147483648 > nbr {
		fmt.Print(int64(nbr))
	} else {
		if nbr < 0 {
			fmt.Print("-")
			nbr *= -1
		}
		i := 0
		nan := ""
		for nbr >= len(str) {
			if nbr >= len(str) {
				nan += string(str[nbr%len(str)])
				nbr = nbr / len(str)
				i++
			}
		}
		nan += string(str[nbr])
		fmt.Print(Reverse(nan))
	}
}

func main() {
	PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
}
