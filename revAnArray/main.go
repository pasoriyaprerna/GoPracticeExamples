package main

import "fmt"

func main() {
	var arrSize int

	fmt.Print("Enter the  Array Size = ")
	fmt.Scan(&arrSize)

	actArr := make([]int, arrSize)
	revArr := make([]int, arrSize)

	fmt.Print("Enter the Even Array Items  = ")
	for i := 0; i < arrSize; i++ {
		fmt.Scan(&actArr[i])
	}

	j := 0

	for i := arrSize - 1; i >= 0; i-- {
		revArr[j] = actArr[i]
		j++
	}

	fmt.Println("\nThe Result of a Reverse Array = ", revArr)
}
