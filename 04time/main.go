package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")
	presentTime := time.Now()
	fmt.Println("Time: ", presentTime)

	// Golang is crazy about these exact number... look for help during setting layout always
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 06 2006 01 02 15 04 05 06 2006 Z070000"))
	createdDate := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local)

	fmt.Println("Generated Date: ", createdDate)
}

func playWithTickers() {

}
