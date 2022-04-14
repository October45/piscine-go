package piscine

func Join(strs []string, sep string) string {
	var result string
	for i, elem := range strs {
		result += elem
		if i != len(strs)-1 {
			result += sep
		}
	}
	return result
}
