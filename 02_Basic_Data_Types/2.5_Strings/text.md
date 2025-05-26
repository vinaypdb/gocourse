# 📘 2.5 Strings in Go

## 🧠 What Are Strings?
Strings in Go are sequences of characters enclosed in double quotes. They are **immutable**, meaning once a string is created, its content cannot be changed.

🧬 **Types of String Literals**
Interpreted Strings: Enclosed in double quotes "". They support escape sequences like `\n`, `\t`, `\"`, etc.

Raw Strings: Enclosed in backticks **`**. These preserve formatting, including newlines and tabs, exactly as typed.

## ✍️ Syntax

```go
var name string = "Alice"       // Interpreted string
message := `Hello,
World!`                         // Raw string literal


```

## ✅ Example Code

```go
package main

import "fmt"

func main() {
    var firstName string = "John"
    lastName := "Doe"
    
    // Concatenation using +
    fullName := firstName + " " + lastName

    // Using raw string literal
    greeting := `Welcome to Go programming!
Learn and explore.`

    fmt.Println("Full Name:", fullName)
    fmt.Println("Greeting Message:")
    fmt.Println(greeting)
}


```

## 🔎 Output

```
Full Name: John Doe
Greeting Message:
Welcome to Go programming!
Learn and explore.


```

## 🔐 Important Points

- Strings are immutable: individual characters cannot be changed.

- You can access characters using indexing (e.g., `s[0]`) but this returns a `byte`.

- Use `len(s)` to get string length.

- You can loop through a string using a `for` loop and `range`.

## 🧹 Summary

| Feature       | Description                                   |
| ------------- | --------------------------------------------- |
| Type          | `string`                                      |
| Literals      | Interpreted (`""`) and Raw (\`\` ` ` \`)      |
| Concatenation | Use `+` to join strings                       |
| Immutability  | Strings can't be changed after creation       |
| Functions     | `len()`, indexing, loops, and string packages |


