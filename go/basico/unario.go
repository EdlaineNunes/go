package main

import "fmt"

func main() {
	x := 1
	y := 2
	x++
	fmt.Printf("valor inicial: %v, valor final %v \n", x, y)
	y-- // y -= 1 ou y = y - 1
	fmt.Println(y)
	// fmt.Println(x == y--)
}
