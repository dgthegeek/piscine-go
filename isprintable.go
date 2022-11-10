package piscine

func IsPrintable(s string) bool {
	bool := true

	for _, value := range s {
		if !(value >= 0 && value <= 15) {
			bool = bool
		} else {
			bool = false
		}
	}
	return bool
}
