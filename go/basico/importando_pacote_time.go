package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Bem-vindo ao Go!")
	fmt.Println("A hora agora é", time.Now().Hour(), ":", time.Now().Minute())
}
