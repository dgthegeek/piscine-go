package piscine

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
	result := 0
	if Prime(nb) {
		return nb
	} else {
		n := nb + 1
		for i := n; i < (n + 10); i++ {
			if Prime(n) {
				result += n
			} else {
				n++
			}
		}
		return n
	}
}
