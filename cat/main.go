package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func reader(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "error"
	}
	return string(content)
}

func print(str string) {
	for _, val := range str {
		z01.PrintRune(val)
	}
}

func main() {
	args := os.Args[1:]

	end := false
	for _, fileName := range args {
		if _, err := os.Stat(fileName); err != nil {
			print("ERROR: open " + fileName + ": no such file or directory\n")
			os.Exit(1)
			return
		}
		print(reader(fileName))
		end = true

	}
	if !end {
		reader := io.TeeReader(os.Stdin, os.Stdout)
		ioutil.ReadAll(reader)
		os.Stdin.Close()
		os.Stdout.Close()
	}
}
