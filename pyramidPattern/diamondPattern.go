package main

import "fmt"

func main() {
	var row int
	fmt.Println("Enter the number of rows")
	fmt.Scanln(&row)

	for i := 1; i <= row; i++ {
		for j := 1; j <= row-1; j++ {
			fmt.Printf("")
		}
		for k := 1; k <= i*2-1; k++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

	for i := row - 1; i > 0; i-- {
		for j := 1; j <= row-1; j++ {
			fmt.Printf("")
		}
		for k := 1; k <= i*2-1; k++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
