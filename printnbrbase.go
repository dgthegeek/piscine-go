package main

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(s int, t string) {
	result := ""
	count := 0
	for _, c := range t {
		if c == c {
			count++
		}
	}
	mx_p := count
	if s < 0 {
		result = "-"
		mx_p *= -1
	}
	if count > 1 {
		for s/mx_p >= count {
			mx_p *= count
		}
		for mx_p != 0 {
			result = result + string(t[s/mx_p])
			s = s - s/mx_p*mx_p
			mx_p /= count
		}
		x := map[rune]bool{}
		for _, c := range t {
			if c == '+' || c == '-' {
				result = "NV"
				break
			}
			if x[c] == false {
				x[c] = true
			} else {
				result = "NV"
				break
			}
		}
	} else {
		result = "NV"
	}
	for _, c := range result {
		z01.PrintRune(c)
	}
}

func main() {
	PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
}
