package piscine

func IterativePower(nb int, power int) int {
	res := nb
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	for p := 2; p <= power; p++ {
		res = res * nb
	}
	return res
}
