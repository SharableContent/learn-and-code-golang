package main

import "fmt"

func main() {
	defer fmt.Println("This is deferred 1 line")
	fmt.Println("Defer the way... ")
	defer fmt.Println("This is deferred 2 line")
	fmt.Println("Come on I am suppose to be the last")
	defer fmt.Println("This is deferred 3 line")
	fmt.Println("I am surely Last one....")
	defer fmt.Println("This is deferred 4 line")

	printInDefer()
}

func printInDefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
