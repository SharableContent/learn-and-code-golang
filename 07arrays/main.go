package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Defining empty list
	var fruitList = [4]string{}
	fmt.Println("Fruit List is", fruitList)
	for i := 0; i < len(fruitList); i++ {
		fruitList[i] = "Fruit-" + strconv.Itoa(i)
	}

	fmt.Println("fruitList: ", fruitList)

	newFruitList := []string{"peach", "apple", "orange", "Green Apples"}
	fmt.Println("fruitList: ", newFruitList)

}
