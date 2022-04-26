package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	table := []int(nil)
	for n > 0 {
		table = append(table, n%10)
		n = n / 10
	}
	SortIntegerTable(table)
	for _, v := range table {
		z01.PrintRune(rune(v + '0'))
	}
}
