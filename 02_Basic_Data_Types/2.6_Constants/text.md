# 📘 2.6 Constants in Go

🧠 **What Are Constants?**

Constants are fixed values that **cannot be changed** during program execution. They help you define meaningful names for values that stay the same throughout your program.


## ✍️ Syntax

```go
const Pi = 3.14         // Implicit type constant
const Greeting string = "Hello, Go!"  // Explicit type constant

```

## ✅ Example Code

```go
package main

import "fmt"

func main() {
    const DaysInWeek = 7
    const Greeting = "Welcome to Go!"

    fmt.Println("Days in a week:", DaysInWeek)
    fmt.Println(Greeting)
}

```

## 🔎 Output

```
Days in a week: 7
Welcome to Go!

```

## 🔐 Important Points

- Constants can be of numeric, string, or boolean types.

- You cannot use `:=` for constants; use `const`.

- Constants are evaluated at compile time.

- Constants improve readability and prevent accidental value changes.

## 🧹 Summary

| Feature     | Description                                   |
| ----------- | --------------------------------------------- |
| Mutability  | Values are fixed and cannot change            |
| Declaration | Use `const` keyword                           |
| Types       | Numeric, string, boolean constants            |
| Use cases   | Fixed values like Pi, configuration, messages |


