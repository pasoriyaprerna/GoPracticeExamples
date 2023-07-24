package main

import (
	"fmt"
)

// create a function
func display(message string) {
	for i := 0; i < 6; i++ {
		fmt.Println(message)
	}
}

func main() {

	// call goroutine
	go display("Process 1")

	display("Process 2")
}
