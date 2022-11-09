package piscine

func Sqrt(nb int) int {
	for f := 1; f*f <= nb; f++ {
		if (f * f) == nb {
			if f%1 == 0 {
				return f
			}
		}
	}
	return 0
}
