package main

import "fmt"

func main() {
	var num, firstNum int
	fmt.Println("Enter the number")
	fmt.Scanln(&num)

	firstNum = num
	for firstNum >= 10 {
		firstNum = firstNum / 10
	}

	fmt.Println("First digit is :", firstNum)
}
