package piscine

func BasicJoin(elems []string) string {
	var result string
	for _, ele := range elems {
		result += string(ele)
	}
	return result
}
