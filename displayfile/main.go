package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := os.Args[1]

	if len(os.Args) < 1 {
		fmt.Println("File name missing")
	} else {
		file, err := ioutil.ReadFile(inputFile)
		if err != nil {
			fmt.Printf("Could not read the file due to this %s error \n", err)
			fileContent := string(file)
			fmt.Println(fileContent)
		}
	}
}
