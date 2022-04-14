package piscine

func BasicJoin(elems []string) string {
	var result string
	for _, elem := range elems {
		result += elem
	}
	return result
}
