package main

import (
	"fmt"
	"log"
)

func sqa(n int) map[int]int {
	var num = make(map[int]int, n)
	for i := 1; i <= n; i++ {
		num[i] = i * i
	}
	return num
}
func main() {
	var n int
	fmt.Println("Enter a number:")
	_, err := fmt.Scanln(&n)
	if err != nil {
		log.Fatal("error", err)
	}
	fmt.Printf("%v", sqa(n))
}
