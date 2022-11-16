package piscine

func SplitWhiteSpaces(s string) []string {

	var list []string
	var mot string
	s = s + " "
	for _, v := range s {
		if string(v) != " " && string(v) != "	" && string(v) != "\n" {
			mot += string(v)
		} else {
			list = append(list, mot)
			mot = ""
		}
	}
	return list
}
