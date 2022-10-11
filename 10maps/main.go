package main

import "fmt"

func main() {
	fmt.Println("Maps in action... ")

	// things inside brackets are Keys and things out side are values
	languages := make(map[string]string)
	fmt.Println("Languages : ", languages)
	fmt.Println("Length Languages : ", len(languages))

	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["node"] = "nodejs"

	fmt.Println("Languages : ", languages)
	fmt.Println("Length Languages : ", len(languages))

	fmt.Println("JS Language : ", languages["js"])

	for key, value := range languages {
		fmt.Printf("key: %s, Value:%s ", key, value)
	}

}
