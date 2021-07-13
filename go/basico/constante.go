package main

import (
	"fmt"
	"reflect"
)

func main() {
	const (
		PI      float32 = 3.14159
		VERSION float32 = 0.1
		B               = 1
	)

	fmt.Println("PI e VERSION são constantes, seus valores são", PI, "e", VERSION)
	fmt.Println("O tipo de B é", reflect.TypeOf(B))
}
