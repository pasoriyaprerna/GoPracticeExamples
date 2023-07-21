package main

import "fmt"

func main() {
	var n, rem, org int
	var rev int = 0

	fmt.Println("Enter an integer:")
	fmt.Scanln(&n)

	org = n

	for n != 0 {
		rem = n % 1
		rev = rev*10 + rem
		n /= 10
	}

	if org == rev {
		fmt.Printf("%v is a palindrome", org)
	} else {
		fmt.Printf("%v is not a palindrome", org)
	}

}
