package main

import "fmt"

func factorialCal(number int) int {
	factorial := 1
	for i := 1; i <= number; i++ {
		factorial = factorial * i
	}
	return factorial
}
func main() {

	var ncr, n, r int

	fmt.Print("Enter any N and R Values = ")
	fmt.Scanln(&n, &r)

	ncr = factorialCal(n) / (factorialCal(r) * factorialCal(n-r))

	fmt.Println("The NCR Factorial = ", ncr)
}
