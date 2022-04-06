package piscine

func Atoi(s string) int {
	if s[0] == '-' && s[1] == '-' {
		return -BasicAtoi(s[1:])
	}
	var i int
	for _, c := range s {
		i = i*10 + int(c-'0')
	}
	return i
}
