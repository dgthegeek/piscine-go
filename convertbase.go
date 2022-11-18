package piscine

func ConvertBase(nbr, baseForm, baseTo string) string {
	result := ""
	my_map := map[rune]int{}

	slice1 := []rune(nbr)
	n1 := len(slice1)

	slice1 := []rune(baseForm)
	n2 := len(slice2)

	slice3 := []rune(baseTo)
	n3 := len(slice3)

	power := 1
	count := 0
	for i := n1 - 1; i >= 0; i-- {
		count = count + my_map[[]rune(nbr)[i]]*power
		power *= n2
	}
	for count != 0 {
		result = string(baseTo[count%n3]) + result
		count /= n3
	}
	return result
}
