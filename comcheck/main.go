package main

import "os"

func main() {
	for _, v := range os.Args {
		if v == "01" || v == "galaxy" || v == "galaxy 01" {
			os.Stdout.Write([]byte("Alert!!!\n"))
			os.Exit(0)
		}
	}
}
