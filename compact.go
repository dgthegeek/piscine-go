package piscine

func Compact(ptr *[]string, length int) int {
	
	cnt := 0
	for i := 0; i < length; i++ {
		if *ptr[i] != " " {
			cnt++
		}
	}
	return cnt
}
