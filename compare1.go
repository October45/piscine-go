package piscine

func Compare1(s1 string, s2 string) int {
	if s1 == s2 {
		return 0
	} else if s1 > s2 {
		return 1
	} else {
		return -1
	}
}
