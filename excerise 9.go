package main

import "fmt"

func main() {
	var i, j, rows, columns int

	var dSumMat [10][10]int

	fmt.Print("Enter the Matrix rows and Columns = ")
	fmt.Scan(&rows, &columns)

	fmt.Println("Enter the Matrix Items to find the Diagonal Sum = ")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Scan(&dSumMat[i][j])
		}
	}
	sum := 0
	for i = 0; i < rows; i++ {
		sum = sum + dSumMat[i][i]
	}
	fmt.Println("The Sum of Matrix Diagonal Elements  = ", sum)
}
