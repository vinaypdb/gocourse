# 🔄  Short Variable Declaration (:=) in Go

## 📘 Explanation

The `:=` operator is used for **short variable declaration** in Go. It allows you to declare and initialize variables without explicitly specifying their type or using the `var` keyword.

Go's compiler **infers the type automatically** from the value on the right-hand side.

## ✍️ Syntax

```go
x := 10            // inferred as int
name := "Gopher"   // inferred as string
flag := true       // inferred as bool
```
You can also declare **multiple variables:**
```
a, b := 100, "Go"

```
## ✅ Example Code

```go
package main

import "fmt"

func main() {
    x := 42
    y := "GoLang"
    z := true

    fmt.Println("x:", x)
    fmt.Println("y:", y)
    fmt.Println("z:", z)
}


```

## 🧪 Sample Output

```
x: 42
y: GoLang
z: true

```

## 🧩 Summary

| Feature             | Description                                |
|---------------------|--------------------------------------------|
| `:=` Operator |Declares and initializes a variable|
| Multiple Assignment       | Multiple variables can be declared at once |
| Scope Limitation    |Only allowed inside functions |
| Cleaner Syntax   |Avoids use of `var` for simple declarations |