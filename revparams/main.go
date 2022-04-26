package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	args_index := len(args) - 1
	for args_index >= 0 {
		for _, str := range args[args_index] {
			z01.PrintRune(str)
		}
		args_index--
		z01.PrintRune('\n')
	}
}
