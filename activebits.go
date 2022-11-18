package piscine

func ActiveBits(n int) int {
	cntt := 0
	remain := 0
	div := 0
	if n < 0 {
		n = -n
	}
	DivMod(n, 2, &div, &remain)
	for div > 0 {
		DivMod(n, 2, &div, &remain)
		n = div
		if remain == 1 {
			cntt++
		}
	}
	return int(cntt)
}
