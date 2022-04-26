package piscine

func AdvancedSortWordArr(table []string, f func(a, b string) int) {
	for i := 0; i < len(table)-1; i++ {
		for j := i + 1; j < len(table); j++ {
			if f(table[i], table[j]) > 0 {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}
