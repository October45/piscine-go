package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 1 && len(os.Args) <= 2 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		defer f.Close()
		b := make([]byte, 100)
		for {
			n, err := f.Read(b)
			if err != nil {
				break
			}
			os.Stdout.Write(b[:n])
		}
	} else if len(os.Args) == 1 {
		fmt.Printf("File name missing\n")
	} else {
		fmt.Printf("Too many arguments\n")
	}
}
