package piscine

func Sqrt(nb int) int {
	var i int = 1
	var n int = 0
	for nb != 0 {
		if i%2 != 0 {
			nb = nb - i
			n++
		}
		if nb < 0 {
			return 0
		}
		i++
	}
	return n
}
