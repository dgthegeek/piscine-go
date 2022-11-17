package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Print()
	} else {
		for _, s := range os.Args[1:] {
			file, err := os.Open(s)
			if err != nil {
				fmt.Println(err.Error())
				break
			} else {
				content, err := ioutil.ReadAll(file)
				if err != nil {
					fmt.Println(err.Error())
					break
				} else {
					fmt.Printf("%s", content)
				}
			}
		}
	}
}
