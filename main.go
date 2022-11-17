/*package main

import (
	"fmt"
	"os"
)

func main() {
	var tableauderune []rune
	var result string
	parametre := []string{"01", "galaxy", "galaxy 01"}

	for i := 1; i < len(os.Args); i++ {
		tableauderune = []rune(os.Args[i])
	}

	for j := 0; j < len(tableauderune); j++ {
		if tableauderune[j] != ' ' {
			result += string(tableauderune[j])
		}
	}

	for _, s := range parametre {
		if result == s {
			fmt.Println("Alert!!!")
		}
	}
}*/
package main

import (
	"fmt"
	"os"
)

func main() {
	for _, v := range os.Args[1:] {
		if v == "01" || v == "galaxy" || v == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
