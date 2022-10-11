package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Number of CPUs in this PC: ", runtime.NumCPU())
	fmt.Println("Operating System of OS: ", runtime.GOOS)
	fmt.Println("Number of goRoutines running here", runtime.NumGoroutine())
}
