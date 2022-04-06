package piscine

func BasicAtoi(s string) int {
	var i int
	for _, c := range s {
		i = i*10 + int(c-'0')
	}
	return i
}
