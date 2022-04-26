package piscine

func AtoiBase(s string, base string) int {
	if len(base) < 2 || Index(base, "+") >= 0 || Index(base, "-") >= 0 {
		return 0
	}
	for i1, char1 := range base {
		for i2, char2 := range base {
			if i1 != i2 && char1 == char2 {
				return 0
			}
		}
	}
	var base_l int = len(base)
	var str_l int = len(s) - 1
	var result int = 0
	for i, c := range s {
		i = str_l - i
		result += Index(base, string(c)) * IterativePower(base_l, i)
	}
	return result
}
