package piscine

func ConcatParams(args []string) string {
	var result string
	n := len(args)
	for i, v := range args {
		if i+1 != n {
			result = result + v + "\n"
		}
	}
	result = result + args[n-1]
	return result
}
