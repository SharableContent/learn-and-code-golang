package main

import "fmt"

func main() {
	var user string
	fmt.Println("[" + user + "]")
	fmt.Printf("Variable is of Type: %T \n", user)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of Type: %T \n", isLoggedIn)

	var smallValue uint8 = 253
	fmt.Println(smallValue)
	fmt.Printf("Variable is of Type: %T \n", smallValue)



	
}
