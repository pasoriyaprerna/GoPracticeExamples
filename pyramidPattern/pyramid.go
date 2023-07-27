package main

import "fmt"

func main() {
	var row int
	fmt.Println("enter the number of rows")
	fmt.Scanln(&row)

	for i := 1; i <= row; i++ {
		for j := 1; j <= row-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k != (2*i - 1); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
