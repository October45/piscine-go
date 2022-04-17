package main

import (
	"os"

	"fmt"
)

func main() {
	if len(os.Args) == 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
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
		os.Stdout.Write("File name missing")
	} else {
		os.Stdout.Write("Too many arguments")
	}
}
