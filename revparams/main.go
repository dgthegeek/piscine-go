package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	var list []string
	var word string
	for _, arg := range args[1:] {
		for i, r := range arg {
			if i >= 0 {
				word += string(r)
			}
		}
		list = append(list, word)
		word = ""

	}

	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}

	for i := 0; i < len(list); i++ {
		mot := list[i]
		for _, v := range mot {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')

	}

	// var list_m [] string
	// for i := len(list); i < 0; i-- {
	// 	list_m = []string(list[i])
	// 	for i := 0; i < len(list_m); i++ {
	// 		z01.PrintRune(list_m[i])
	// 		z01.PrintRune('\n')

	// 	}
	// }
}
