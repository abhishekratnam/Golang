package main

import (
	"fmt"
)

type user struct {
	userName   string
	userAge    int
	userGender string
}

func main() {
	var u user

	u.userName = "Abhishek"
	u.userAge = 21
	u.userGender = "M"

	fmt.Println("Name of the user is", u.userName)
	fmt.Println("Age of the user is", u.userAge)
	fmt.Println("Gender of the user is", u.userGender)

}
