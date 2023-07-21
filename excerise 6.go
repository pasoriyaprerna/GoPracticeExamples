package main

import "fmt"

func main() {
	var num, i, j int

	var identMat [10][10]int

	fmt.Print("Enter the Matrix Size = ")
	fmt.Scan(&num)

	fmt.Print("Enter the Matrix Items = ")
	for i = 0; i < num; i++ {
		for j = 0; j < num; j++ {
			fmt.Scan(&identMat[i][j])
		}
	}
	flag := 1
	for i = 0; i < num; i++ {
		for j = 0; j < num; j++ {
			if identMat[i][j] != 1 && identMat[j][i] != 0 {
				flag = 0
				break
			}
		}
	}
	if flag == 1 {
		fmt.Println("It is an Idenetity Matrix")
	} else {
		fmt.Println("It is Not an Idenetity Matrix")
	}
}
