package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(addTraditionally(100, 200))
	fmt.Println(performOperation(Add, 100, 200, 300, 400, 500, 600, 700))
	fmt.Println(performOperation(Multiply, 100, 200, 300, 400, 500, 600, 700))
}

type Operation int

const (
	Add Operation = iota
	Multiply
)

// traditional, define params..
func addTraditionally(num1 int32, num2 int32) int32 {
	return num1 + num2
}

// traditional, define params..
func performOperation(operation Operation, numbers ...int64) int64 {
	switch operation {
	case Add:
		return add(numbers)
	case Multiply:
		return multiply(numbers)
	}
	return 0
}

func add(numbers []int64) int64 {
	var result int64 = 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func multiply(numbers []int64) int64 {
	var result int64 = 0
	for _, number := range numbers {
		result *= number
	}
	return result
}

func squares(numbers []int64) []int64 {
	var result []int64 = make([]int64, len(numbers))
	for index, number := range numbers {
		result[index] = number * number
	}
	return result
}

func sqrt(numbers ...float64) []float64 {
	var result []float64 = make([]float64, len(numbers))
	for index, number := range numbers {
		result[index] = math.Sqrt(number)
	}
	return result
}
