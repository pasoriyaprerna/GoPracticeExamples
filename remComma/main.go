package main

import (
	"fmt"
	"strconv"
	"strings"
)

func remove(input string) []int {
	num := strings.Split(input, ",")
	len := len(num)
	var n = make([]int, len)
	for i, v := range num {
		s := strings.Trim(v, " ")
		n[i], _ = strconv.Atoi(s)
	}
	return n
}

func main() {
	result := remove("10, 200,20,50,60,70")
	fmt.Println(result)

}
