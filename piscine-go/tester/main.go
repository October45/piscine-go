package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.ConvertBase("125", "0123456789", "0123456789"))
	fmt.Println(piscine.ConvertBase("101011", "01", "0123456789"))
	fmt.Println(piscine.ConvertBase("7D", "0123456789ABCDEF", "0123456789"))
	fmt.Println(piscine.ConvertBase("uoi", "01", "0123456789"))
	fmt.Println(piscine.ConvertBase("ci", "choumi", "0123456789"))
	fmt.Println(piscine.ConvertBase("101011", "01", "0123456789"))
	fmt.Println(piscine.ConvertBase("13303", "0123456789", "0123456789ABCDEF"))
	fmt.Println(piscine.ConvertBase("931287", "0123456789", "0123456789ABCDEF"))
	fmt.Println(piscine.ConvertBase("bcaccacaaaacc", "abc", "01"))
}
