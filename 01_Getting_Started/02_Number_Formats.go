package main

import "fmt"

func main() {
	num := 42

	fmt.Printf("%d \t %b \t %#X \n", num, num, num)
	fmt.Printf("%d - %b - %x \n", num, num, num)
	fmt.Printf("%d - %b - %#x \n", num, num, num)
	fmt.Printf("%d - %b - %#X \n", num, num, num)
}

