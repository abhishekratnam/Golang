package main

import "fmt"

var c, python, java bool
var i, j int = 1, 2

func main() {
	var i int
	fmt.Println(i, c, python, java)
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

}
