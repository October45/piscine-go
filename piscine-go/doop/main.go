package main

import (
	"os"
)

func Contains(str string, substr string) bool {
	for _, c := range substr {
		for _, sc := range str {
			if c == sc {
				return true
			}
		}
	}
	return false
}

func Atoi(str string) int {
	var sum int
	var negative bool = false
	for i, c := range str {
		if c == '-' {
			negative = true
			continue
		}
		sum += int(c) - '0'
		if len(str) > 1 && i != len(str)-1 {
			sum *= 10
		}
	}
	if negative {
		sum *= -1
	}
	return sum
}

func RevString(str string) string {
	var rev string
	for i := len(str) - 1; i >= 0; i-- {
		rev += string(str[i])
	}
	return rev
}

func RevAtoi(nbr int) string {
	var str string
	if len(str) == 0 {
		for nbr > 0 {
			str += string(rune(nbr%10) + '0')
			nbr /= 10
		}
	}
	if len(str) == 0 {
		nbr *= -1
		for nbr > 0 {
			str += string(rune(nbr%10) + '0')
			nbr /= 10
		}
		if len(str) == 0 {
			str = "0"
		} else {
			str += "-"
		}
	}
	return RevString(str)
}

func main() {
	args := os.Args
	if len(args) != 4 {
		os.Exit(0)
	}
	args = args[1:]
	if !Contains(args[1], "+-*/%") {
		os.Exit(0)
	}
	if (args[1] == "%" && args[2] == "0") || (args[1] == "/" && args[2] == "0") {
		switch args[1] {
		case "%":
			os.Stdout.Write([]byte("No modulo by 0\n"))
			os.Exit(0)
		case "/":
			os.Stdout.Write([]byte("No division by 0\n"))
			os.Exit(0)
		}
	}
	for i := range args[0] {
		if args[0][i] >= 'a' && args[0][i] <= 'z' {
			os.Exit(0)
		}
		if args[0][i] >= 'A' && args[0][i] <= 'Z' {
			os.Exit(0)
		}
	}
	for i := range args[2] {
		if args[2][i] >= 'a' && args[2][i] <= 'z' {
			os.Exit(0)
		}
		if args[2][i] >= 'A' && args[2][i] <= 'Z' {
			os.Exit(0)
		}
	}
	num1 := Atoi(args[0])
	num2 := Atoi(args[2])
	if args[0] >= "9223372036854775807" || args[2] >= "9223372036854775807" {
		os.Exit(0)
	}
	var sum int
	switch args[1] {
	case "+":
		sum = num1 + num2
		if num1 > 0 && num2 > 0 && sum < 0 {
			os.Exit(0)
		}
	case "-":
		sum = num1 - num2
		if num1 < 0 && num2 < 0 && sum > 0 {
			os.Exit(0)
		}
	case "*":
		sum = num1 * num2
		if num1 > 0 && num2 > 0 && sum < 0 {
			os.Exit(0)
		}
	case "/":
		sum = num1 / num2
	case "%":
		sum = num1 % num2
	}
	os.Stdout.Write([]byte(RevAtoi(sum) + "\n"))
}
