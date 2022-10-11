package main

import (
	"fmt"
)

type Student struct {
	name  string
	email string
}

func main() {

	//ifElseFunc()
	forLoopFunc()
}

func ifElseFunc() {
	fmt.Println("IfElse session... ")

	loginCount := 23
	var result string

	if loginCount > 10 {
		result = "Regular User"
	} else {
		result = "New user"
	}

	fmt.Println("Result: ", result)

	// Variable Initialiazing and condition on same line
	if num := 3; num < 10 {
		fmt.Println("Initialized and checked on same line")
	}

	if loginCount > 10 && loginCount%2 == 0 {
		fmt.Println("Multiple conditions")
	}
}

func forLoopFunc() {

	days := []string{"Monday", "Tuesday", "wednesday", "Thursday", "Friday"}

	//Traditional
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	//Range Loop
	for index, data := range days {
		fmt.Printf("Data: %s, Index in array: %d \n", data, index)
	}

}
