package main

import "fmt"

func main() {
	rows := 5 // Change this value to adjust the size of the diamond pattern

	// Upper half of the diamond pattern
	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print(i)
		}
		fmt.Println()
	}

	// Lower half of the diamond pattern
	for i := rows - 1; i >= 1; i-- {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print(i)
		}
		fmt.Println()
	}
}
