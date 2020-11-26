package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	su := 1
	for su < 1000 {
		su += su
	}
	fmt.Println(su)
	s := 1
	for s < 1000 {
		s += s
	}
	fmt.Println(s)
}
