package main

import "fmt"

func main() {
	var lista = [...]string{"a", "b", "c"}
	a := "Teste"
	b := a
	fmt.Printf("Valor de a = %v ponteiro de a = %v \n", a, &a)
	fmt.Printf("Valor de b = %v ponteiro de b = %v \n", b, &b)
	fmt.Println(lista, &lista, &lista[0])
}
