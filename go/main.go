package main

import (
	"fmt"
	"time"
)

func greet(prefix, name string) {
	fmt.Printf("%s %s!\n", prefix, name)
}

func formatGreet(prefix, name string) string {
	return prefix + " " + name + "! \n"
}

func main() {
	fmt.Printf("Hello World!\n")
	fmt.Printf("Time now: %s\n", time.Now())
	greet("Hello", "Edlaine")
	greet("Hello", "World!")
	fmt.Printf(formatGreet("Imprimir", "formatado!"))
	// fmt.Print(str.Reverse("Hello"))

	// enMsg := "Hello"
	// ptMsg := "Olá"
	// name := "Edlaine"

	// fmt.Println(str.Swap(enMsg, name))
	// second, first := str.SwapTwo(ptMsg, name)
	// fmt.Println(second, first)

	var str string
	var integer int
	var boolean bool
	var char rune = 'a'

	fmt.Printf("Zero value %v, type: %T\n", str, str)
	fmt.Printf("Zero value %v, type: %T\n", integer, integer)
	fmt.Printf("Zero value %v, type: %T\n", boolean, boolean)
	fmt.Printf("Zero value %v, type: %T\n", char, char)

}
