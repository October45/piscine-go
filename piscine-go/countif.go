package piscine

func CountIf(f func(string) bool, a []string) int {
	count := 0
	for _, v := range a {
		if f(v) == true {
			count++
		}
	}
	return count
}
