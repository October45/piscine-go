package piscine

func Map(f func(int) bool, a []int) []bool {
	b := make([]bool, len(a))
	for i, v := range a {
		b[i] = f(v)
	}
	return b
}
