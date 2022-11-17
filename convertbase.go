package piscine

func ConvertBase(nbr, baseForm, baseTo string) string {
	var result string
	nbrlist := []rune(nbr)
	n1 := len(nbrlist)

	baseformlist := []rune(baseForm)
	n2 := len(baseformlist)

	baseTolist := []rune(baseTo)
	n3 := len(baseTolist)

	st_map := map[rune]int{}

	power := 1
	count := 0
	for i := n1 - 1; i >= 0; i-- {
		count = count + st_map[[]rune(nbr)[i]]*power
		power *= n2
	}
	for count != 0 {
		result = string(baseTo[count%n3]) + result
		count /= n3
	}
	return result
}
