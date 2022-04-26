package piscine

func Max(nbrs []int) int {
	if len(nbrs) == 0 {
		return 0
	}
	max := nbrs[0]
	for _, n := range nbrs {
		if n > max {
			max = n
		}
	}
	return max
}
