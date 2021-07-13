package main

// import(
// 	"errors"
// )

func Reverse(s string) string {
	// if s == " "{
	// 	return null
	// }
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Swap(first, second string) string {
	return second + first
}

func SwapTwo(first, second string) (string, string) {
	return second, first
}
