package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a char to check digit:")
	r, _, _ := reader.ReadRune()
	if unicode.IsDigit(r) {
		fmt.Printf("%c is a digit", r)
	} else {
		fmt.Printf("%c is not a digit", r)
	}
}
