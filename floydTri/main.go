package main

import "fmt"

func main() {
	var row int
	num := 1
	fmt.Println("enter total number of rows:")
	fmt.Scanln(&row)
	fmt.Println("------------------")
	for i := 1; i <= row; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d", num)
			num++
		}
		fmt.Println()
	}
}
