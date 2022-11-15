package piscine

func SplitWhiteSpaces(s string) []string {
	var tab []string
	splite := ""
	for i, str := range s {
		if str > 32 {
			splite += string(str)
		} else if splite != "" {
			tab = append(tab, splite)
			splite = ""
		}
		if i == len(s)-1 {
			tab = append(tab, splite)
		}
	}
	return tab
}

// This is my first attemp to build my fonction but it ignore the last word

// package piscine

// import "fmt"

// func SplitWhiteSpaces(s string) []string {
// 	var result []string
// 	slice := []rune(s)
// 	var list [][]rune
// 	var wrd []rune
// 	for _, v := range slice {
// 		if v != ' ' {
// 			wrd = append(wrd, v)
// 		} else {
// 			list = append(list, wrd)
// 			wrd = nil
// 		}
// 	}
// 	for _, v := range list {
// 		s := string(v)
// 		result = append(result, s)
// 	}

// 	return result
// }
