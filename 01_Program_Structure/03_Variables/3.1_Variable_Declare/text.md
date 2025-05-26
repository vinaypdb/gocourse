# 🔄   What is a Variable in Go?

## 📘 Explanation

A **variable** in Go is a **named storage location** used to hold a value that your program can read and change. Variables are typed, meaning each variable must have a specific data type (e.g., `int`, `string`, `bool`, etc.).

## ✍️ Syntax

```go
var name type            // Declaration without value
var age int = 25         // Declaration with value
x := "GoLang"            // Short variable declaration (inside functions only)

```
## ✅ Example Code

```go
package main

import "fmt"

func main() {
    var name string = "Alice"    // var with type and value
    var age int                  // var with type only (gets zero value)
    city := "New York"           // short variable declaration with type inference

    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("City:", city)
}


```

## 🧪 Sample Output

```
Name: Alice
Age: 0
City: New York


```

📝 Key Points

- Use `var` for **explicit type declaration** or global variables.

- Use `:=` for **type inference** inside functions.

- Uninitialized variables get **zero values** by default.