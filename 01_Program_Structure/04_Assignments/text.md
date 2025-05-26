# 🔄 Assignments in Go

## 📘 Explanation

An **assignment** in Go means giving a value to a declared variable. Once a variable is declared, you use the assignment operator (`=`) to assign or update its value.

## ✍️ Syntax 

Go uses the `var`, `const`, and `type` keywords for declarations.

```go
variableName = value         // Single assignment
a, b = value1, value2        // Multiple assignment
a, b = b, a                  // Swap values

```

## ✅ Example Code

```go
package main

import "fmt"

func main() {
    var x int
    x = 10 // Assignment after declaration

    y := 20 // Shorthand declaration + assignment

    // Multiple assignment
    a, b := 1, 2

    // Swapping values
    a, b = b, a

    fmt.Println("x:", x)
    fmt.Println("y:", y)
    fmt.Println("a:", a)
    fmt.Println("b:", b)
}

```

## 🧪 Sample Output

```
x: 10
y: 20
a: 2
b: 1

```

## 🧩 Key Points:

| Feature               | Description                                   |
| --------------------- | --------------------------------------------- |
| Single Assignment     | Assigns one value to a variable               |
| Multiple Assignment   | Allows assigning multiple values in one line  |
| Swapping Values       | A Go idiom: `a, b = b, a`                     |
| Reassignment          | You can reassign new values after declaration |
| Shorthand Declaration | Uses `:=` to declare and assign in one line   |
