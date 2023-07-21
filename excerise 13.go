package main

import (
	"fmt"
	"unicode"
)

func main() {
	character := 'r' // The character you want to check

	if unicode.IsUpper(character) {
		fmt.Printf("The character '%c' is an uppercase letter.\n", character)
	} else {
		fmt.Printf("The character '%c' is not an uppercase letter.\n", character)
	}
}
