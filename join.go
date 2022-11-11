package piscine

func Join(strs []string, sep string) string {
	var result string
	slice := []string(strs)
	for i := 0; i < len(slice)-1; i++ {
		result += slice[i] + string(sep)
	}
	result += slice[len(slice)-1]
	return result
}
