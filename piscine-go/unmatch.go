package piscine

func Unmatch(nbrs []int) int {
	for i := 0; i < len(nbrs)-1; i++ {
		for j := i + 1; j < len(nbrs); j++ {
			if nbrs[i] > nbrs[j] {
				nbrs[i], nbrs[j] = nbrs[j], nbrs[i]
			}
		}
	}
	for 0 < len(nbrs) {
		if len(nbrs)-1 != 0 {
			if nbrs[0] == nbrs[1] {
				nbrs = nbrs[2:]
			} else {
				return nbrs[0]
			}
		} else {
			return nbrs[0]
		}
	}
	return -1
}
