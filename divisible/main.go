package main

import (
	"fmt"
	"strconv"
	"strings"
)

func divisible(less, more int) string {
	var num []string
	for i := less; i <= more; i++ {
		if i%7 == 0 && i%5 != 0 {
			num = append(num, strconv.Itoa(i))
		}
	}
	return strings.Join(num, ",")
}
func main() {
	result := divisible(10, 200)
	fmt.Println(result)
}
