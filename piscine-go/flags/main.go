package main

import (
	"fmt"
	"os"
)

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

func printHelp() {
	fmt.Printf("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n")
	os.Exit(0)
}

func parseFlags(args []string) []string {
	var flags []string
	for i, arg := range args {
		if arg == "--help" || arg == "-h" {
			printHelp()
		}
		if len(arg) != 0 {
			if arg[0] == '-' {
				flags = append(flags, arg)
				args = append(args[:i], args[i:]...)
			}
		}
	}
	return flags
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
	}
	args := os.Args[1:]
	flags := parseFlags(args)
	var print_s string
	var sort bool = false
	for _, flag := range flags {
		if Index(flag, "-i") != -1 {
			print_s += flag[Index(flag, "=")+1:]
		}
		if Index(flag, "-o") != -1 {
			sort = true
		}
	}
	print_s = args[len(args)-1] + print_s
	if !sort {
		fmt.Println(print_s)
	} else {
		var runes []rune
		for _, r := range print_s {
			runes = append(runes, r)
		}
		for i := 0; i < len(runes)-1; i++ {
			for j := i + 1; j < len(runes); j++ {
				if runes[i] > runes[j] {
					runes[i], runes[j] = runes[j], runes[i]
				}
			}
		}
		fmt.Println(string(runes))
	}
}
