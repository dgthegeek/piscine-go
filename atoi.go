package piscine

func Atoi(s string) int {
	sign := 1
	slice := []rune(s)
	n := len(s)
	var transformed int

	for i := 0; i < n; i++ {
		if n == 0 {
			return 0
		}
		if slice[i] < '0' || slice[i] > '9' {
			if slice[i] < '0' {
				if slice[i] == '+' {
				} else {
					sign = -1
				}
				// '>1' so i zont get the error 'idex out of range with a certain lenght'
				if n > 1 {
					if slice[i+1] == '-' || slice[i+1] == '+' {
						return 0
					}
				}
				if slice[i] == ' ' {
					return 0
				}
			} else {
				return 0
			}
		} else {
			transformed *= 10
			transformed += int(slice[i]) - '0'
		}
	}
	final := transformed * sign
	return final
}
