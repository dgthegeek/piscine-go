package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		var tab []int = nil
		return tab
	} else {
		tab := []int{}

		for m := min; m < max; m++ {
			tab = append(tab, m)
		}
		return tab
	}
}
