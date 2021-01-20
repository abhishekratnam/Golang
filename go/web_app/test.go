// substring_matching.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Make HTTP GET request
	fmt.Println("Enter Web URL: ")
	// var then variable name then variable type
	var URL string
	// Taking input from user
	fmt.Scanln(&URL)
	response, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	// Get the response body as a string
	dataInBytes, err := ioutil.ReadAll(response.Body)
	pageContent := string(dataInBytes)
	// Find a substr
	var subStr string
	fmt.Println("Enter the Keyword you wanna find: ")
	fmt.Scanln(&subStr)
	hasSubstr := strings.Contains(pageContent, subStr)
	if hasSubstr == false {
		fmt.Println("Try another web pal !!!!")
		os.Exit(0)
	} else {
		fmt.Println("Gotcha!!..web contains that keyword punk")
	}
}
