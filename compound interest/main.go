package main

import (
	"fmt"
	"math"
)

func main() {
	var Pamt, IntRate, timePeriod, comp, ci float64
	fmt.Print("Enter principal amount :")
	fmt.Scanln(&Pamt)

	fmt.Print("Enter Interest Rate :")
	fmt.Scanln(&IntRate)

	fmt.Print("Enter Number of years :")
	fmt.Scanln(&timePeriod)

	ci = Pamt * (math.Pow((1 + IntRate/100), timePeriod))

	comp = ci - Pamt
	fmt.Println("Compond Interest :", comp)
	fmt.Println("Future Compound Interest :", ci)
}
