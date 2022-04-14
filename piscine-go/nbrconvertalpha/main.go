package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PerfAtoi(s string) int {
	var i int = 0
	for _, c := range s {
		if !(c >= '0' && c <= '9') {
			return 0
		}
		i = i*10 + int(c-'0')
	}
	return i
}

func Atoi(s string) int {
	var i int = 0
	if len(s) == 0 || s == "-" || s == "+" {
		return i
	}
	for _, c := range s {
		if !((c >= '0' && c <= '9') || c == '-' || c == '+') {
			return i
		}
	}
	if s[0] == '-' && s[1] == '-' {
		return i
	}
	if s[0] == '-' && len(s) != 1 {
		s = s[1:]
		i = PerfAtoi(s)
		i = -i
	} else if s[0] == '+' && len(s) != 1 {
		s = s[1:]
		i = PerfAtoi(s)
	} else {
		i = PerfAtoi(s)
	}
	return i
}

func main() {
	args := os.Args[1:]
	caps := false
	if len(args) == 0 {
		return
	}
	if args[0] == "--upper" {
		caps = true
	}
	for i, arg := range args {
		if i == 0 && caps {
			continue
		}
		char_int := Atoi(arg)
		if char_int == 0 || char_int > 26 {
			z01.PrintRune(' ')
		} else {
			if caps {
				z01.PrintRune(rune(char_int + 'A' - 1))
			} else {
				z01.PrintRune(rune(char_int + 'a' - 1))
			}
		}
	}
	z01.PrintRune('\n')
}
