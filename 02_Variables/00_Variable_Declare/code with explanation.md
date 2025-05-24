# 🧠 Variable Declaration and Assignment in Go

## 📄 Code
```go
package main

import "fmt"

func main() {
	var message string
	message = "Hello World."
	fmt.Println(message)
}
```
🧪 Output

```
Hello World.
```

## 📦 Package and Import

- package main: Defines the entry point of a Go executable program.

- import "fmt": Imports the fmt package, which provides I/O formatting functions such as Println.

## 🧮 Variable Declaration and Assignment

- var message string
Declares a variable named message with an explicit type string. Since it's not initialized, it takes the zero value, which is "" (empty string).

- message = "Hello World."
Assigns the string "Hello World." to the previously declared message variable.

## 🖨️ Printing the Value

- fmt.Println(message)
Prints the value of the message variable to the standard output. In this case, it prints:
Hello World.

## 📌 Summary

- This example shows how to declare a string variable using var with an explicit type.

- The variable is assigned a value after declaration.

- fmt.Println is used to print the variable’s content to the terminal.
