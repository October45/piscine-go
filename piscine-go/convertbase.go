package piscine

func NumFromBase(num int, base string) string {
	mod := len(base)
	i := 0
	var str string
	if num == 0 {
		str = "0"
		return str
	}
	for j := 1; j <= num%mod; j++ {
		i++
	}
	for j := -1; j >= num%mod; j-- {
		i++
	}
	if num/mod != 0 {
		PrintNumBase(num/mod, base)
	}
	str += string(rune(base[i]))
	return str
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	n := AtoiBase(nbr, baseFrom)
	return NumFromBase(n, baseTo)
}
