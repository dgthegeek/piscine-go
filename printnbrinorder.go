package main

import "github.com/01-edu/z01"

func main() {

	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var tab [10]int
	for n != 0 {
		tab[n%10]++
		n /= 10
	}
	for i := 0; i < 10; i++ {
		for tab[i] > 0 {
			z01.PrintRune(rune(i) + '0')
			z01.PrintRune(rune(i))
			tab[i]--

		}
	}
	z01.PrintRune(tab)
}


	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}
