# 🔄 5. Multiple Variable Declaration and Assignment in Go

## 📘 Explanation

Go allows you to declare and assign multiple variables **in a single line** using either:

- The `var` keyword (with explicit types)

- The short declaration operator `:=` (with type inference)

This improves code **readability** and **reduces redundancy.**

## ✍️ Syntax

```go
var a, b int = 10, 20  // using var keyword
x, y := "Go", 2025     // using short declaratio
```

## ✅ Example

```go
package main

import "fmt"

func main() {
    var a, b int = 10, 20        // explicit type declaration
    x, y := "Go", 2025           // short declaration with type inference

    fmt.Println("a:", a, "b:", b)
    fmt.Println("x:", x, "y:", y)
}

```

## 🧪 Sample Output

```
a: 10 b: 20
x: Go y: 2025


```

## 🧩 Summary

| Feature             | Description                                |
|---------------------|--------------------------------------------|
| Multiple Declaration|Declare multiple variables on one line
       |
| `var` Keyword           | 	Explicitly specify types |
| `:=` Short Declaration     | Infers types based on the assigned values  |
| Cleaner Code   | Reduces repetition and improves readability  |