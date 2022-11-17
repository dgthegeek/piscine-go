package main

import "fmt"

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}

	pas := 0

	for start != 1 {
		if start%2 == 0 {
			start = (start / 2)
		} else {
			start = 3*start + 1
		}
		pas++
	}

	return pas
}

func main() {
	steps := CollatzCountdown(12)
	fmt.Println(steps)
}
