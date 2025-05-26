# 📘 2.4 Booleans in Go

## 🧠 What Are Booleans?

Booleans represent truth values in Go. They can only be `true` or `false`. Booleans are primarily used in conditions, loops, and logical operations to control the flow of a program.

## ✍️ Syntax

```go
var isActive bool = true    // Declaring a boolean variable
isReady := false            // Short variable declaration

```

## ✅ Example Code

```go
package main

import "fmt"

func main() {
    var isLoggedIn bool = true   // User is logged in
    isAdmin := false             // User is not admin

    fmt.Println("Is Logged In?", isLoggedIn)
    fmt.Println("Is Admin?", isAdmin)

    if isLoggedIn && !isAdmin {
        fmt.Println("Regular user access")
    } else if isAdmin {
        fmt.Println("Admin access")
    } else {
        fmt.Println("Guest access")
    }
}

```

## 🔎 Output

```
Is Logged In? true
Is Admin? false
Regular user access

```

## 🧹 Summary

| Feature        | Description                         |
| -------------- | ----------------------------------- |
| Boolean Values | Can only be `true` or `false`       | 
| Usage          | Used in conditions and logic checks 
| Declaration    | Use `var` or `:=`                   |
| Logical Ops    | Supports `&&`, \`                   | 


