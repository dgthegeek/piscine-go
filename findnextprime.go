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
	nexprime := 0

	for i := nb; i <= nb+1; i++ {
		if !Prime(i) {
			nb++
		} else {
			nexprime += i
			break
		}
	}

	return nexprime
}
