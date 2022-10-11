package main

import "fmt"

func main() {
	fmt.Println("We will play with Pointers")
	var numberValue int = 24
	numberPtr := &numberValue
	fmt.Println("Value of number: ", numberPtr)
	fmt.Println("Value of pointer number: ", *numberPtr)
	fmt.Printf("Type of number: %T", numberPtr)
	*numberPtr = *numberPtr * 100
	fmt.Println("Value of number: ", numberPtr)
	fmt.Println("* Value of number: ", *numberPtr)
	fmt.Println("& Value of number: ", &numberPtr) // Not sure what the use of this here.  Reference of a pointer, which is pointer to something
	fmt.Printf("Type of number: %T", numberPtr)

	var nextNumPtr *int = &numberValue
	fmt.Println("Value of nextNumPtr: ", nextNumPtr)
	fmt.Println("* Value of nextNumPtr: ", *nextNumPtr)

}
