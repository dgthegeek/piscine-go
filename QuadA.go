package main

import (
	"fmt"
)

func printSymbol(x int) {
	fmt.Printf("o")
	if x > 2 {
		for i := 0; i < x-2; i++ {
			fmt.Printf("-")
		}
	}
}

func quad(x, y int) {
	if x == 1 && y == 1 {
		fmt.Printf("o")
	}

	if x == 1 && y > 1 {
		fmt.Println("o")
		if y > 2 {
			for i := 0; i < y-2; i++ {
				fmt.Println("|")
			}
		}
		fmt.Println("o")
	}

	if x > 1 && y == 1 {
		printSymbol(x)
		fmt.Printf("o")
	}

	if x > 1 && y > 1 {
		fmt.Printf("o")
		if x > 2 {
			for i := 0; i < x-2; i++ {
				fmt.Printf("-")
			}
		}

		fmt.Println("o")

		if y > 2 {
			for i := 0; i < y-2; i++ {
				fmt.Printf("|")
				if x > 2 {
					for i := 0; i < x-2; i++ {
						fmt.Printf(" ")
					}
				}
				fmt.Println("|")

			}
		}

		printSymbol(x)
		fmt.Printf("o")

	}
	fmt.Println(" ")
}

func main() {
	quad(20, 1)
}
