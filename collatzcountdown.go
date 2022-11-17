package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}

	result := 1

	for start != 1 {
		if start%2 == 0 {
			result = start / 2
		} else {
			result = 3*start + 1
		}
		result++
	}

	return result
}
