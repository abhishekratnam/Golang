package main

import (
	"fmt"
	"unicode"
)

func Password(pass string) bool {
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}

	if !upp || !low || !num || !sym || tot < 8 {
		return false
	}

	return true
}

func main() {
	for _, s := range []string{
		"noob",
		"noobPassword",
		"noobPa##word",
		"b3tterPa$$w0rd",
	} {
		fmt.Printf("Password(%q) = %v\n", s, Password(s))
	}
}
