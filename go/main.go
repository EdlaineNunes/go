package main

import (
	"fmt"
	"time"
)

func greet(prefix, name string){
	fmt.Printf("%s %s!\n", prefix, name)
}

func formatGreet(prefix, name string) string{
	return prefix + " " + name + "! \n"
}


func main() {
	fmt.Printf("Hello World!\n")
	fmt.Printf("Time now: %s\n", time.Now())
	greet("Hello", "Edlaine")
	greet("Hello", "World!")
	fmt.Printf(formatGreet("Imprimir", "formatado!"))
	// fmt.Printf(str.Reverse("Hello"))
}

