package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
     for a := '0'; a <= '9'; a = a + 1 {
		for b := '0'; b <= '9'; b = b + 1 {
			i := b + 1
			for j := a; j <= '9'; j++ {
				for ; i <= '9'; i++ {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(' ')
					z01.PrintRune(j)
					z01.PrintRune(i)
					if b < '8' || a < '9' || i < '9' || j < '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				i = '0'
			}
		}
	}
	z01.PrintRune('\n')
	}