package main

import "fmt"
import "strconv"

func main() {
	b, err := strconv.ParseBool("true")
	f, err := strconv.ParseFloat("3.1415", 64)
	i, err := strconv.ParseInt("-42", 10, 64)
	u, err := strconv.ParseUint("42", 10, 64)

	fmt.Printf("%v [%T], %v [%T], %v [%T], %v [%T]\n", b, b, f, f, i, i, u, u)
	fmt.Printf("%v [%T]\n", err, err)
}
