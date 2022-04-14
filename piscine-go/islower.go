package piscine

func IsLower(s string) bool {
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			continue
		} else {
			return false
		}
	}
	return true
}
