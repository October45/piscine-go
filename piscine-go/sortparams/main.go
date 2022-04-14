package main

import (
	"os"

	"github.com/01-edu/z01"
)

func SortTable(table []string) {
	for i := 0; i < len(table)-1; i++ {
		for j := i + 1; j < len(table); j++ {
			if table[i] > table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}

func main() {
	args := os.Args[1:]
	SortTable(args)
	for _, arg := range args {
		for _, b := range arg {
			z01.PrintRune(rune(b))
		}
		z01.PrintRune('\n')
	}
}
