package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	var k []int
	fmt.Println(k, len(k), cap(k))
	if k == nil {
		fmt.Println("nil!")
	}
}
