# 🔄    Identifiers in Go

## 📘 Explanation

An **identifier** is the name used to identify variables, functions, constants, types, packages, and labels in Go programs. Identifiers must follow specific rules and conventions to be valid and meaningful in the Go language.

## ✍️ Syntax Rules for Identifiers:

- Must start with a **letter** (A-Z, a-z) or an underscore _.

- Can contain **letters, digits** (0-9), and **underscores** after the first character.

- Are **case-sensitive** (`count` and `Count` are different identifiers).

- Cannot be a Go keyword (like `var`, `func`, `if`).

- Identifiers starting with an **uppercase letter** are **exported** (visible outside the package).

- Identifiers starting with a **lowercase letter** are **unexported** (package-private).

## ✅ Example Code

```go
package main

import "fmt"

func main() {
    // Valid identifiers
    var name string = "Alice"
    age := 30

    // Identifiers are case-sensitive
    var Name string = "Bob"

    fmt.Println(name)  // prints: Alice
    fmt.Println(age)   // prints: 30
    fmt.Println(Name)  // prints: Bob
}

```

## 🧪 Sample Output

```
Alice
30
Bob
```
## 🧩 Explanation:

- `name`, `age`, and `Name` are identifiers.

- Notice `name` and `Name` are different due to case sensitivity.

- Identifiers start with letters and can contain letters and digits.

- This program shows only valid identifiers in use.

## 🧩 📘 Reserved Keywords in Go 
Go has a fixed set of **25 reserved keywords** that have special meaning in the language. These **cannot be used as identifiers** (i.e., you can't name a variable or function using them).

```
break      default       func       interface   select
case       defer         go         map         struct
chan       else          goto       package     switch
const      fallthrough   if         range       type
continue   for           import     return      var

```


## 🧩 Key Points:

| Feature             | Description                                |
|---------------------|--------------------------------------------|
| Start Character |Letter (A-Z, a-z) or underscore `_`|
| Subsequent Characters   | Letters, digits (0-9), or underscore `_` |
| Case Sensitivity   |Identifiers are case-sensitive |
| Exported Identifiers  |Start with uppercase letter (accessible outside package)|
| Unexported Identifiers |Start with lowercase letter (accessible only inside package)|
| No Keywords Allowed |Cannot use reserved Go keywords as identifiers|