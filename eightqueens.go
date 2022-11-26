package piscine

import "github.com/01-edu/z01"

const N = 8

var position = [N]int{}

func PrintNbr99(n int) {
	t := 1
	if n < 0 {
		t = -1
		z01.PrintRune('-')
	}
	if n != 0 {
		f := (n / 10) * t
		if f != 0 {
			PrintNbr99(f)
		}
		k := (n % 10 * t) + '0'
		z01.PrintRune(rune(k))
	} else {
		z01.PrintRune('0')
	}
}

func talivre(queennumber, rowposition int) bool {
	for i := 0; i < queennumber; i++ {
		other_row_pos := position[i]

		if other_row_pos == rowposition ||
			other_row_pos == rowposition-(queennumber-i) ||
			other_row_pos == rowposition+(queennumber-i) {
			return false
		}
	}
	return true
}

func resolverpuzzle(k int) {
	if k == N {
		for i := 0; i < N; i++ {
			PrintNbr99(position[i] + 1)
		}
		z01.PrintRune('\n')
	} else {
		for i := 0; i < N; i++ {
			if talivre(k, i) {
				position[k] = i
				resolverpuzzle(k + 1)
			}
		}
	}
}

func EightQueens() {
	resolverpuzzle(0)
}
