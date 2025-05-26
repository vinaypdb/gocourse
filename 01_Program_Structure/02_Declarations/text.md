# 🔄  Declarations in Go

## 📘 Explanation

**Declarations** in Go are used to **introduce names** into a program—these names can be for variables, constants, types, functions, or packages. A declaration tells the compiler about the name and what kind of entity it represents.

## ✍️ Syntax 

Go uses the `var`, `const`, and `type` keywords for declarations.

```go
var variableName type       // Variable declaration
const constantName = value  // Constant declaration
type TypeName struct { ... } // Type declaration

```

## ✅ Example Code

```go
package main

import "fmt"

// Global variable declaration
var globalVar int = 100

// Constant declaration
const Pi = 3.14

// Type declaration
type Person struct {
    name string
    age  int
}

func main() {
    // Local variable declaration
    var x int = 10
    y := 20 // Short variable declaration

    // Using declared values
    fmt.Println("x:", x)
    fmt.Println("y:", y)
    fmt.Println("Pi:", Pi)

    // Using declared type
    p := Person{name: "Alice", age: 25}
    fmt.Println("Person:", p)
}


```

## 🧪 Sample Output

```
x: 10
y: 20
Pi: 3.14
Person: {Alice 25}

```

## 🧩 Key Points:

| | Declaration Type | Keyword | Description                                    |
| ---------------- | ------- | ------------------------------------------------ |
| Variable         | `var`   | Declares a named variable with a type            |
| Constant         | `const` | Declares a name with a fixed, unchangeable value |
| Type             | `type`  | Declares a new named type (e.g., struct)         |
| Short Variable   | `:=`    | Used inside functions for quick declaration      |
| Multiple         |         | You can declare multiple names in one line       |
