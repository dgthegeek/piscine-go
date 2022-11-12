package main

import (
	"os"

	"github.com/01-edu/z01"
)

func SortIntegerTable2(table []string) {
	i := 1
	for i < len(table) {
		if table[i-1] > table[i] {
			tmp := table[i]
			table[i] = table[i-1]
			table[i-1] = tmp
			i = 1
		} else {
			i++
		}
	}
}

func main() {
	arg := os.Args[2:]
	var list []string
	for i := 0; i < len(arg); i++ {
		list = append(list, arg[i])
	}
	SortIntegerTable2(list)
	for i := 0; i < len(list); i++ {
		a := list[i]
		for j := 0; j < len(a); j++ {
			z01.PrintRune(rune(a[j]))
		}
		z01.PrintRune('\n')
	}
}
