package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for _, str := range args {
		for _, ch := range str {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
