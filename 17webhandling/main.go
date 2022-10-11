package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const testUrl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	//readFromTheWebURL(testUrl)
	parseURLAndPrint(testUrl)
}

func readFromTheWebURL(urlToRead string) {
	fmt.Println("Learning the Web Requests now")

	response, err := http.Get(urlToRead)
	handleNilError(err)

	fmt.Printf("Response of type: %T\n", response)
	defer response.Body.Close() //Callers responsibility to close this connection

	dataBytes, err := ioutil.ReadAll(response.Body)
	content := string(dataBytes)
	handleNilError(err)

	fmt.Println("Data received from Response: ", content)
}

func parseURLAndPrint(urlToRead string) {

	result, err := url.Parse(urlToRead)
	handleNilError(err)

	fmt.Println("result.Scheme: ", result.Scheme)
	fmt.Println("result.Host: ", result.Host)
	fmt.Println("result.Path: ", result.Path)
	fmt.Println("result.Port: ", result.Port())
	fmt.Println("result.RawQuery: ", result.RawQuery)

	queryParams := result.Query()

	fmt.Printf("The type of query params are: %T\n", queryParams)

	for key, value := range queryParams {
		fmt.Printf("Key: %s, Value: %s \n", key, value)
	}
}

func handleNilError(err error) {
	if err != nil {
		panic(err)
	}
}
