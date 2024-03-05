package main

import (
	"fmt"
)

var a, b, c int

func main() {
	// message := "Hello World!"
	// var message string
	// message = "Hello World!"
	// var a rune = 'a'

	// fmt.Println(a)
	// fmt.Println(message)
	a, b, c = 1, 2, 3
	a, _ = 1, 6
	fmt.Println(a, b, c)
	print()
}

func print() {
	// a, b, c := 4, 5, 6
	fmt.Println(a, b, c)
}
