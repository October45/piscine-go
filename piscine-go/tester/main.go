package main

import (
	"fmt"

	"piscine"
)

func PrintList(l *piscine.List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

func main() {
	link := &piscine.List{}
	link2 := &piscine.List{}

	fmt.Println("----normal state----")
	piscine.ListPushBack(link2, 1)
	PrintList(link2)
	piscine.ListRemoveIf(link2, 1)
	fmt.Println("------answer-----")
	PrintList(link2)
	fmt.Println()

	fmt.Println("----normal state----")
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "Hello")
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "There")
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "How")
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "are")
	piscine.ListPushBack(link, "you")
	piscine.ListPushBack(link, 1)
	PrintList(link)

	piscine.ListRemoveIf(link, 1)
	fmt.Println("------answer-----")
	PrintList(link)
}
