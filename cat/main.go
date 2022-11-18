package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		reader := io.TeeReader(os.Stdin, os.Stdout)
		ioutil.ReadAll(reader)
		os.Stdin.Close()
		os.Stdout.Close()
	} else {
		for _, arg := range arguments {
			ReadFile(string(arg))
		}
	}
}

func ReadFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		for _, val := range "ERROR: " + err.Error() {
			z01.PrintRune(val)
		}
		z01.PrintRune('\n')
		os.Exit(1)
	}
	for _, dt := range data {
		z01.PrintRune(rune(dt))
	}
}
