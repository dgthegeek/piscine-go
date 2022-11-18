package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	n := len(os.Args)
	for i := 1; i < n; i++ {
		content, err := ioutil.ReadFile(os.Args[i])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(content))
		fmt.Println()
	}
}
