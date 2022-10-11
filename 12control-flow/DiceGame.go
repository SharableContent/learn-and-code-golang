package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func PlayDiceGame() {

}

func switchCaseFunc() {

	// Seeded to create number
	rand.Seed(time.Now().UnixNano())

	for {
		fmt.Println(" Do you want to continue Playing ? Press any key to continue, followed by enter...  ")
		fmt.Println("Press 'q' to quit the game")
		inputFromUser := getByteInputFromUser()
		if inputFromUser == "q" {
			fmt.Println("Sorry to see you leave the Game... ")
		}

		fmt.Println("You want to continue, great....")
		fmt.Println("Guess Next Spot")
		inputFromUser = getByteInputFromUser()

		// Add timer
		addTimer()
		//diceNumber := rand.Intn(6)

		// ToDo: Complete the Game

	}
}
func getNextSpotFromUser() int32 {
	var inputReader = bufio.NewReader(os.Stdin)
	var inputFromUser int
	// in case of a byte
	inputBuffer, _, _ := inputReader.ReadLine()
	inputFromUser, _ = strconv.ParseInt(bytes.NewBuffer(inputBuffer).String())

	return inputFromUser
}

func getByteInputFromUser() string {
	var inputReader = bufio.NewReader(os.Stdin)
	var inputFromUser string
	// in case of a byte
	inputBuffer, _ := inputReader.ReadByte()
	// In case of string
	// inputFromUser = bytes.NewBuffer(inputBuffer).String()
	inputFromUser = fmt.Sprintf("%s", inputBuffer)
	return inputFromUser
}

func addTimer() {
	fmt.Println("here you go....\n Guess the next number")
	for i := 0; i < 5; i++ {
		time.Sleep(500)
		fmt.Print(".")
	}
}
