package main

import "fmt"

func main() {
	var row, col int
	var matFirst, matSec [10][10]int

	fmt.Println("enter the number of rows and columns")
	fmt.Scan(&row, &col)
	fmt.Println("Enter the items for first matrix")
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Scan(&matFirst[i][j])
		}
	}

	fmt.Println("Enter the items for second matrix")
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Scan(&matSec[i][j])
		}
	}

	if matFirst == matSec {
		fmt.Println("Both matrix are equal")
	} else {
		fmt.Println("Both matrix are unequal")
	}
}
