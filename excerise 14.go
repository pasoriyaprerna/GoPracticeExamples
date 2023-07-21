package main

import "fmt"

func main() {

	var cublen, cubWidth, cubHeight, cubSA, cubVolume, cubLA float32

	fmt.Print("Enter the Length of a Cuboid = ")
	fmt.Scanln(&cublen)
	fmt.Print("Enter the Width of a Cuboid  = ")
	fmt.Scanln(&cubWidth)
	fmt.Print("Enter the Height of a Cuboid = ")
	fmt.Scanln(&cubHeight)

	cubSA = 2 * (cublen*cubWidth + cublen*cubHeight + cubWidth*cubHeight)
	cubVolume = cublen * cubWidth * cubHeight
	cubLA = 2 * cubHeight * (cublen + cubWidth)

	fmt.Println("\nThe Volume of a Cuboid         = ", cubVolume)
	fmt.Println("The Surface Area of a Cuboid     = ", cubSA)
	fmt.Println("Lateral Surface Area of a Cuboid = ", cubLA)
}
