package piscine

import (
	"fmt"
	"runtime"
)

func printRuntime() string {
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)
	fmt.Println("NumCPU:", runtime.NumCPU())
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())
	return runtime.Version()
}
