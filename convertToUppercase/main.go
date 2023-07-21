package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter char to convert into Uppercase:")
	lwch, _, _ := reader.ReadRune()

	if unicode.IsLetter(lwch) {
		up := unicode.ToUpper(lwch)
		fmt.Printf("Uppercase Char of %c = %c", lwch, up)
	} else {
		fmt.Printf("Enter valid char")
	}

}
