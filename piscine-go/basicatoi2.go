package piscine

func BasicAtoi2(s string) int {
	var i int
	for _, c := range s {
		if !(c >= '0' && c <= '9') {
			return 0
		}
		i = i*10 + int(c-'0')
	}
	return i
}
