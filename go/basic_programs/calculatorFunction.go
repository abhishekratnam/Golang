package main

import "fmt"

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}
func div(x, y int) int {
	return x / y
}
func mul(x, y int) int {
	return x * y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(sub(42, 13))
	fmt.Println(mul(42, 13))
	fmt.Println(div(42, 13))
}
