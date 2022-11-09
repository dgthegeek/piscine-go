package main

import (
	"fmt"
)

func Prime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	nexprime := 0

	for i := nb; i <= nb+1; i++ {
		if Prime(i) == false {
			nb++
		} else {
			nexprime += i
			break
		}
	}

	return nexprime
}

func main() {
	fmt.Println(FindNextPrime(8))
	fmt.Println(FindNextPrime(10))
}
