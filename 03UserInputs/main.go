package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	mulitiplierTest()
}

func userInputAndParsing() {
	inputReader := bufio.NewReader(os.Stdin)
	//fmt.Println("Please provide us feedback for our cooking....")
	//line, _ := inputReader.ReadString('\n')
	//
	//fmt.Printf("Line input as: %s", line)
	//fmt.Printf("Type of the input is: %T", line)

	fmt.Println("Please provide us feedback for our cooking between 1 to 5.")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Input entered by User is : ", numRating)
}

func mulitiplierTest() {
	inputReader := bufio.NewReader(os.Stdin)
	var total float64 = 1
	fmt.Println("I can multiply as long as you want, keep entering numbers and once done, press q")

	for {
		fmt.Println("Type next number and press enter... type 'q' to quit ")
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		if strings.TrimSpace(input) != "q" {
			newNumber, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
			fmt.Println("You have entered: ", newNumber)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}

			total = total * newNumber
		} else {
			fmt.Printf("Looks like you are tired now, the total is : %f", total)
			break
		}

	}

}
