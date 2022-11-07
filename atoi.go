package piscine

import "strconv"

func Atoi(s string) int {
            a := 0
           b, e := strconv.Atoi(s)
	if e == nil{
		 a = b 
	}

return a
}