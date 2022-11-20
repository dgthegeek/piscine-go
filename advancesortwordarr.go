package piscine

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	quickSrot2(array, 0, len(array)-1)
}
