package piscine

func Rot14(s string) string {
	var result string
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			if c+14 > 'Z' {
				result += string(c + 14 - 26)
			} else {
				result += string(c + 14)
			}
		} else if c >= 'a' && c <= 'z' {
			if c+14 > 'z' {
				result += string(c + 14 - 26)
			} else {
				result += string(c + 14)
			}
		} else {
			result += string(c)
		}
	}
	return result
}
