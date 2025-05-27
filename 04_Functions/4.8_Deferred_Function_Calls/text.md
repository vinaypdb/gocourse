# 5.5. Deferred Function Calls in Go

---

## 📌 What is `defer` in Go?

The `defer` keyword in Go is used to **postpone the execution** of a function until **just before the surrounding function returns**.

Deferred calls are **executed in LIFO (Last In, First Out)** order.

---

## ❓ Why Use `defer`?

* ✅ For **cleanup activities** like closing files, unlocking mutexes, or closing database connections.
* ✅ Ensures that important steps are **executed no matter how the function exits**.
* ✅ Helps improve **code readability and maintainability**.

---

## 🧾 Syntax of `defer`

```go
defer functionCall()
```

* The function is **scheduled to run later**, not run immediately.
* Arguments to the deferred function are **evaluated immediately**, but the call is delayed.

---

## ✅ Example 1: Basic Defer Usage

```go
package main

import "fmt"

func main() {
    defer fmt.Println("World")
    fmt.Println("Hello")
}
```

### 🧪 Output

```
Hello
World
```

---

## ✅ Example 2: Multiple Defers (LIFO Order)

```go
package main

import "fmt"

func main() {
    defer fmt.Println("First")
    defer fmt.Println("Second")
    defer fmt.Println("Third")
    fmt.Println("Start")
}
```

### 🧪 Output

```
Start
Third
Second
First
```

---

## ✅ Example 3: File Closing with Defer

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("example.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close() // ensures the file is closed

    file.WriteString("Hello, file!")
    fmt.Println("Write complete")
}
```

---

## ✅ Example 4: Deferred Function with Arguments

```go
package main

import "fmt"

func main() {
    a := 5
    defer fmt.Println("Deferred value of a:", a)
    a = 10
    fmt.Println("Changed value of a:", a)
}
```

### 🧪 Output

```
Changed value of a: 10
Deferred value of a: 5
```

🔎 The value `5` was **captured** at the moment `defer` was called.

---

## 🧷 Key Points Recap

| Concept                 | Description                           |
| ----------------------- | ------------------------------------- |
| `defer`                 | Postpones function call until return  |
| LIFO order              | Last deferred call runs first         |
| Immediate argument eval | Arguments are evaluated at defer time |
| Used for cleanup        | Common for `Close`, `Unlock`, etc.    |

---

## 🧭 Summary

* `defer` is used to **schedule function calls** to run after the current function completes.
* Great for **cleaning up resources** in a predictable and readable way.
* Deferred calls are executed in **reverse order** (LIFO).
* Arguments are **evaluated when defer is called**, not when executed.

> 🛠️ Use `defer` to write robust and maintainable code, especially when managing resources like files, locks, or network connections.
