package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println("Slice Operations")
	//fruitList := []string{"peach", "apple", "orange", "Green Apples"}
	//fmt.Println("fruitList: ", newFruitList)
	//fmt.Printf("Type of fruitList is: %T\n", newFruitList)
	//newFruitList := append(fruitList, "banana", "creame")
	//fmt.Println("fruitList: ", newFruitList)
	//
	//newFruitList = append(fruitList[2:])  // start with index 2 and go till end
	//newFruitList = append(fruitList[:3])  // start with index 0 and count 3 items
	//newFruitList = append(fruitList[1:3]) // start with index 1 and count 3 elments e.g. 1, 2, 3,  Indexes
	//
	//fmt.Println("fruitList: ", newFruitList)

	highScores := make([]int, 4)
	highScoresPtr := &highScores
	fmt.Println("highscores: ", highScores)

	highScores[0] = 10
	highScores[2] = 16
	highScores[1] = 20
	highScores[3] = 31
	fmt.Println("highscores: ", highScores)
	fmt.Println("Length of highscores: ", len(highScores))
	fmt.Println("Index of highscores: ", highScoresPtr)

	// throws index out of bound
	// highScores[4] = 777

	// But appending works
	highScores = append(highScores, 666, 777)
	fmt.Println("highscores: ", highScores)
	sort.Ints(highScores)
	fmt.Println("highscores: ", highScores)

	// remove using Append
	indexToRemove := 2
	highScores = append(highScores[:indexToRemove], highScores[indexToRemove+1:]...)
	fmt.Println("highscores: ", highScores)
}
