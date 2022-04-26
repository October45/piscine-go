package piscine

func Abort(a, b, c, d, e int) int {
	int_a := []int{a, b, c, d, e}
	for i := 0; i < len(int_a)-1; i++ {
		for j := i + 1; j < len(int_a); j++ {
			if int_a[i] > int_a[j] {
				int_a[i], int_a[j] = int_a[j], int_a[i]
			}
		}
	}
	return int_a[2]
}
