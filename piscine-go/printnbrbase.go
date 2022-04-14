package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNumBase(num int, base string) {
	mod := len(base)
	i := 0
	if num == 0 {
		z01.PrintRune('0')
		return
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
	z01.PrintRune(rune(base[i]))
	return
}

func PrintNbrBase(n int, base string) {
	if len(base) < 2 || Index(base, "+") >= 0 || Index(base, "-") >= 0 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	for i1, char1 := range base {
		for i2, char2 := range base {
			if i1 != i2 && char1 == char2 {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}
	if n < 0 {
		z01.PrintRune('-')
	}
	PrintNumBase(n, base)
}
