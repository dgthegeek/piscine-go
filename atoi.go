package piscine

func BasicAtoi(s string) int {
	slice := []rune(s)
	n := len(s)
	var transformed int
	sign := 1
	for i := 0; i < n; i++ {
		if slice[i] < '0' || slice[i] > '9' {
			return 0
		} else {

			if slice[i] == '+' || slice[i] == '-' {
				if slice[i+1] == '+' || slice[i+1] == '-' {
					return 0
				} else {
					if slice[i] == '+' {
						sign = 1
					} else if slice[i] == '-' {
						sign = -1
					}
				}
			}

			transformed *= 10
			transformed += int(slice[i]) - '0'
		}
	}
	return transformed * sign
}
