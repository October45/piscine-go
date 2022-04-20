package main

import (
	"fmt"
	"os"
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
	c := args[1]
	args = args[2:]
	didWoopsie := false
	for i := 0; i < len(args); i++ {
		f, err := os.Open(args[i])
		if err != nil {
			fmt.Printf(err.Error() + "\n")
			didWoopsie = true
		} else {
			defer f.Close()
			b, err := os.ReadFile(args[i])
			if err != nil {
				fmt.Printf(err.Error() + "\n")
				didWoopsie = true
				continue
			}
			if Atoi(c) < len(b) {
				b = b[len(b)-Atoi(c):]
			}
			if i != 0 {
				fmt.Printf("\n")
			}
			fmt.Printf("==> " + args[i] + " <==\n")
			fmt.Printf(string(b))
		}
	}
	if didWoopsie {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
