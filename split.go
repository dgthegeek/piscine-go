package piscine

func Split(s, sep string) []string {
	lists := []rune(s)
	n1 := len(lists)
	listsep := []rune(sep)
	n2 := len(listsep)
	for i := 0; i < n1-n2; i++ {
		if sep == s[i:n2+i] {
			s = s[:i] + " " + s[n2+i:]
			n1 = n1 - n2
		}
	}
	return SplitWhiteSpaces(s)

}