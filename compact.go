package piscine

func Compact(ptr *[]string) int {
	res := 0
	var arr []string
	for _, v := range *ptr {
		if v != "" {
			arr = append(arr, v)
			res++
		}
	}
	*ptr = arr
	return res
}
