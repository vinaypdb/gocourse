package main

import "fmt"

func main() {
	num := 42

	fmt.Printf("%d \t %b \t %#X \n", num, num, num)
	fmt.Printf("%d - %b - %x \n", num, num, num)
	fmt.Printf("%d - %b - %#x \n", num, num, num)
	fmt.Printf("%d - %b - %#X \n", num, num, num)
}

/*
Walkthrough of syntax and logic:

package main
Declares the package. main means this is an executable program with an entry point.

import "fmt"
Imports the fmt package for formatted I/O functions like Printf.

func main() { ... }
The main function is the entry point of every Go program.

num := 42
Declares and initializes an integer variable num with the value 42.

fmt.Printf
Prints formatted output using verbs inside the format string:

%d — prints decimal (base 10) integer.

%b — prints binary representation.

%x — prints hexadecimal in lowercase.

%X — prints hexadecimal in uppercase.

%#x or %#X — prints hex with 0x (or 0X) prefix.

\t inserts a tab space between columns for neat alignment.

\n adds a newline after each print. */
