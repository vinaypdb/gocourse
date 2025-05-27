# 🔁 switch Statement in Go

## 📘 Explanation

**📌 What is a switch Statement?**

A `switch` statement provides a **cleaner and more readable** way to handle **multiple conditional branches** compared to chained `if-else` statements. Go's switch is flexible and allows:

- Expression-less `switch`

- Multiple case values

- `fallthrough` behavior

- Type switches

**🤔 Why Use switch?**

Use it when:

- You have many conditions on the same variable.

- You want **clear**, **concise**, and **efficient** control flow.

- You prefer less boilerplate compared to `if-else`.


## 🧱 Syntax

```go
switch expression {
case value1:
    // block
case value2:
    // block
default:
    // optional fallback block
}

```

## ✅ Example: Switch with Expression

```go
package main

import "fmt"

func main() {
    day := 3

    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    default:
        fmt.Println("Unknown day")
    }
}


```

## 🧪 Sample Output

```go
Wednesday

```

## ✅ Example: Multiple Values per Case

```go
switch day {
case 1, 2, 3, 4, 5:
    fmt.Println("Weekday")
case 6, 7:
    fmt.Println("Weekend")
default:
    fmt.Println("Invalid day")
}

```

## ✅ Example: Switch Without Expression

```go
x := 15

switch {
case x < 10:
    fmt.Println("Less than 10")
case x < 20:
    fmt.Println("Less than 20")
default:
    fmt.Println("20 or more")
}

```
This behaves like an if-else chain.

## ✅ Example: Using fallthrough

```go
switch x := 2; x {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
    fallthrough
case 3:
    fmt.Println("Three") // Will be executed due to fallthrough
}


```
`fallthrough` forces the next case to execute **even if the condition doesn’t match**.

## ❌ Common Mistakes Beginners Make

**1. Expecting fallthrough by default**

Unlike C/C++, Go **does NOT fall through** unless explicitly stated.

```go
switch 2 {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
case 3:
    fmt.Println("Three") // ❌ Will not run unless fallthrough is used
}

```
✅ Remember: `break` exits the whole loop.

**2. Using complex logic directly inside case**

```go
switch x {
case x > 10: // ❌ Compile-time error
    fmt.Println("Greater than 10")
}

```
✅ Fix:

```go
switch {
case x > 10:
    fmt.Println("Greater than 10")
}

```


## 💡 Key Points

| Point                                              | Description                        |
| -------------------------------------------------- | ---------------------------------- |
| `switch` is a cleaner alternative to `if-else`     | Especially for multiple conditions |
| Go switches do not fall through by default         | Use `fallthrough` explicitly       |
| You can use `switch` with or without an expression | Enables flexibility                |
| `default` is optional                              | Executes if no other case matches  |


## 🧠 Summary

- Go's `switch` is flexible and cleaner than `if-else`.

- Supports expression-based and condition-based usage.

- Avoids fallthrough unless specified.

- Great for decision trees, state machines, and branching logic.

