package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	my_slice := []rune(s)
	for i := 0; i < len(my_slice); i++ {
		z01.PrintRune(my_slice[i])
	}
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		content, err := ioutil.ReadFile(os.Args[i])
		if err != nil {
			mes := "ERROR: " + err.Error() + "no such file or directory"
			PrintStr(mes)
			z01.PrintRune('\n')
		}
		PrintStr(string(content))
	}
}
