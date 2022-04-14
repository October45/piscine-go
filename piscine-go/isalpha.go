package piscine

func IsAlpha(s string) bool {
	if len(s) == 0 {
		return true
	}
	for _, c := range s {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			continue
		} else {
			return false
		}
	}
	return true
}
