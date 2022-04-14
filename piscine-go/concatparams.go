package piscine

func ConcatParams(args []string) string {
	var result string
	for i := range args {
		result += args[i]
		if i != len(args)-1 {
			result += "\n"
		}
	}
	return result
}
