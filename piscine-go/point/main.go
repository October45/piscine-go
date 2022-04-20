package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func NumFromBase(num int, base string) string {
	mod := len(base)
	i := 0
	if num == 0 {
		return "0"
	}
	for j := 1; j <= num%mod; j++ {
		i++
	}
	for j := -1; j >= num%mod; j-- {
		i++
	}
	return NumFromBase(num/mod, base) + string(base[i])
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)

	result := "x = " + NumFromBase(points.x, "0123456789")[1:] + ", y = " + NumFromBase(points.y, "0123456789")[1:] + "\n"

	for _, c := range result {
		z01.PrintRune(c)
	}
}
