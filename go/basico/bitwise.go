package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a, b int64 = 11, 10

	e := strconv.FormatInt(a&b, 2)
	ou := strconv.FormatInt(a|b, 2)
	xor := strconv.FormatInt(a^b, 2)

	fmt.Printf("E => %v %v \n", a&b, e)     // 11 & 10 = 10
	fmt.Printf("Ou => %v %v \n", a|b, ou)   // 11 | 10 = 11
	fmt.Printf("Xor => %v %v \n", a^b, xor) // 11 ^ 10 = 01
}
