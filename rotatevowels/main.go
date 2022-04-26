package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Contains(str string, ch rune) bool {
	for _, c := range str {
		if c == ch {
			return true
		}
	}
	return false
}

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	if len(s) == 0 {
		return -1
	}
	for i := 0; i < len(s); i++ {
		if s[i] == toFind[0] {
			if len(toFind) == 1 {
				return i
			}
			if len(s[i:]) >= len(toFind) {
				if s[i:i+len(toFind)] == toFind {
					return i
				}
			}
		}
	}
	return -1
}

func main() {
	if !(len(os.Args) >= 2) {
		z01.PrintRune('\n')
		os.Exit(0)
	}
	args := os.Args[1:]
	var args_s []rune
	for _, arg := range args {
		for _, c := range arg {
			args_s = append(args_s, c)
		}
		args_s = append(args_s, ' ')
	}
	var vowel_pos_rev string
	var vowel_pos string
	var vowels string = "aeiouAEIOU"
	for i, c := range args_s {
		if Contains(vowels, c) {
			vowel_pos += string(rune(i))
			vowel_pos_rev = string(rune(i)) + vowel_pos_rev
		}
	}
	if len(vowel_pos) == 0 {
		for _, c := range args_s {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
		os.Exit(0)
	}
	for i := 0; i < len(vowel_pos)/2; i++ {
		args_s[vowel_pos[i]], args_s[vowel_pos_rev[i]] = args_s[vowel_pos_rev[i]], args_s[vowel_pos[i]]
	}
	for _, c := range args_s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
