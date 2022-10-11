package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const fileName = "./my_go_file.txt"

func main() {
	fmt.Println("File Handling... ")
	writeFile()
	defer readFile()
}

func writeFile() {
	content := "This need to go in file, save it well..."
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Printf("We have write %d long string in file", length)
	defer file.Close()
}

func readFile() {
	//var dataBytes []byte
	dataBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text Data inside file is \n ", dataBytes)
	fmt.Println("Text Data inside file is \n ", string(dataBytes))

}
