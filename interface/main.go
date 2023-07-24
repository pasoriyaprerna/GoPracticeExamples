package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func (s *square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t *triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

//access method of interface
func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	t := triangle{base: 10, height: 10}
	s := square{sideLength: 5}
	//struct implements the interface
	printArea(&t)
	printArea(&s)
}

// create a function
// func display(message string) {

// 	fmt.Println(message)
// }

// func main() {

// 	// call goroutine
// 	go display("Process 1")

// 	display("Process 2")
// }
