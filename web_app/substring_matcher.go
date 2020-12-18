// substring_matching.go
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func trashFinder(URL, word string) bool {
	response, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	// Get the response body as a string
	dataInBytes, err := ioutil.ReadAll(response.Body)
	pageContent := string(dataInBytes)
	// Find a substr
	hasSubstr := strings.Contains(pageContent, word)
	if hasSubstr {
		return true
	} else {
		return false
	}
}
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
func main() {
	var keyword string
	fmt.Println("Enter the keyword to match")
	fmt.Scanln(&keyword)
	// Make HTTP GET request
	file, err := os.Open("file.txt")
	// var then variable name then variable type
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()
	var matchedUrls []string
	for _, eachline := range txtlines {
		if trashFinder(eachline, keyword) {
			matchedUrls = append(matchedUrls, eachline)
		}
	}
	fmt.Println(matchedUrls)

}
