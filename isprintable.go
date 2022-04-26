package piscine

func IsPrintable(str string) bool {
	for _, c := range str {
		if c >= ' ' && c <= '~' {
			continue
		} else {
			return false
		}
	}
	return true
}
