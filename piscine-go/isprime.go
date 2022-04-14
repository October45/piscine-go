package piscine

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	if nb == 2 || nb == 3 || nb == 5 || nb == 7 || nb == 11 {
		return true
	} else if nb < 11 {
		return false
	}
	for i := 2; i < nb/2; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
