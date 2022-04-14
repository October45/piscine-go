package piscine

func IsUpper(s string) bool {
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			continue
		} else {
			return false
		}
	}
	return true
}
