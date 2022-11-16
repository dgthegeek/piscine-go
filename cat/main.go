package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		content, err := ioutil.ReadFile(os.Args[i])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf(string(content))
	}
}
