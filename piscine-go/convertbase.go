package piscine

func NumFromBase(num int, base string) string {
	mod := len(base)
	i := 0
	if num == 0 {
		return "0"
	}
	for j := 1; j <= num%mod; j++ {
		i++
	}
	for j := -1; j >= num%mod; j-- {
		i++
	}
	return NumFromBase(num/mod, base) + string(base[i])
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	n := AtoiBase(nbr, baseFrom)
	return NumFromBase(n, baseTo)[1:]
}
