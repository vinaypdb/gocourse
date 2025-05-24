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

- `package main`: Declares this file as part of the main package, which makes it a standalone executable.
- `import "fmt"`: Brings in Go’s standard library `fmt` package for formatted input/output operations.

---

## 🔧 Function: main()

- `func main()`: This is the program’s entry point. Go looks for `main()` when starting execution.

---

## 🧮 Variable Declaration

- `var message string`: Declares a string variable named `message`. Since it’s not initialized, it holds the zero value: an empty string (`""`).

---

## ✍️ Assignment

- `message = "Hello World."`: Assigns the string `"Hello World."` to the `message` variable.

---

## 🖨️ Output Formatting

- `fmt.Println(message)`: Prints the content of `message` to the terminal.

---

## 📌 Summary

- Demonstrates Go’s **two-step variable declaration and assignment**.
- Helps in scenarios where a variable is declared upfront and its value is decided later.