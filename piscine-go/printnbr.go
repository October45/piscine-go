package piscine

import "github.com/01-edu/z01"

func PrintNum(num int) {
	i := '0'
	if num == 0 {
		z01.PrintRune('0')
		return
	}
	for j := 1; j <= num%10; j++ {
		i++
	}
	for j := -1; j >= num%10; j-- {
		i++
	}
	if num/10 != 0 {
		PrintNum(num / 10)
	}
	z01.PrintRune(i)
	return
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	PrintNum(n)
}
