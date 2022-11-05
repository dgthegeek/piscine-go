package piscine

func PrintNbr(n int) {
	i := 1
	if n < 0 {
		i = -1
		z01.PrintRune('-')
	}
	if n != 0 {
		new := (n / 10) * i
		if f != 0 {
			PrintNbr(new)
		}
		j := (n % 10 * i) + '0'
		z01.PrintRune(rune(j))
	}else {
		z01.PrintRune('0')
	}
}