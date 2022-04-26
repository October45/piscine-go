package piscine

func Split(s, sep string) []string {
	for i := 0; i < len(s); i++ {
		if Index(s, sep) != -1 {
			s = s[0:Index(s, sep)] + " " + s[Index(s, sep)+len(sep):]
		}
	}
	return SplitWhiteSpaces(s)
}
