package piscine

func ActiveBits(n int) uint {
	cnt := 0
	for n > 0 {
		cnt += n & 1
		n >>= 1
	}
	return uint(cnt)
}
