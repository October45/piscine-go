package piscine

func IsNumeric(str string) bool {
	for _, c := range str {
		if c >= '0' && c <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
