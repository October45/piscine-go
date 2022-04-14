package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	}
	if IsPrime(nb) {
		return nb
	}
	var i int = nb + 1
	for nb < i {
		if IsPrime(i) {
			return i
		} else {
			i++
		}
	}
	return 0
}
