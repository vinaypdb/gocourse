# 🔄 5. Type Conversion in Go

## 📘 Explanation

In Go, a **constant** is a value that cannot be changed once it is assigned. Constants are declared using the `const` keyword and must be assigned a value at the time of declaration.

- Constants can be of type: string, boolean, or numeric (int, float, etc.).

- Constants **cannot** be declared using `:=` shorthand.

- You **cannot assign the result of a function call** to a constant — the value must be known at compile time.

## ✍️ Syntax

```go
const name = value
```
## ✍️ Or with a type:

```go
const pi float64 = 3.14
```

## ✅ Example

```go
package main

import "fmt"

func main() {
    const appName = "GoTutorial"
    const version = 1.0
    const isReleased = true

    fmt.Println("App Name:", appName)
    fmt.Println("Version:", version)
    fmt.Println("Released:", isReleased)

    // appName = "NewApp" // ❌ This will cause a compile-time error
}

```

## 🧪 Sample Output

```
App Name: GoTutorial
Version: 1
Released: true
```

## 🧩 Summary

| Feature             | Description                                |
|---------------------|--------------------------------------------|
| Immutable | Constant values can't be changed after set         |
| Declared with              | 	`const` keyword                         |
| Compile-time      | Value must be known at compile time   |
| Types    | Can be int, float, string, bool  |