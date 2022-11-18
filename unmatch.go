package piscine

func Unmatch(arr []int) int {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				arr[i] = -1
				arr[j] = -1
				break
			}
		}
	}
	for k := 0; k < len(arr); k++ {
		if arr[k] != -1 {

			return arr[k]
		}
	}
	return -1
}
