package main

import "fmt"

func main() {
	a := 10
	b := 20

	fmt.Println("Before swap:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// Swap values
	a, b = b, a

	fmt.Println("After swap:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
