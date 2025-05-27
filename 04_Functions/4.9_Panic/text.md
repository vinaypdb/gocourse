# 5.6. Panic in Go

---

## 📌 What is `panic` in Go?

A `panic` in Go is a **built-in function** that **stops the normal execution** of the current goroutine.

When a function panics:

* It starts **unwinding the stack** (running deferred functions).
* Then the program **crashes** with a log message.

---

## ❓ Why Use `panic`?

* ✅ To signal a **critical error** that the program **cannot recover from**.
* ✅ Common in **library code** when contract is violated (e.g., index out of range).
* ✅ Often used in combination with `recover` to handle unexpected situations gracefully.

> ⚠️ Use `panic` sparingly. Prefer error handling for expected conditions.

---

## 🧾 Syntax of `panic`

```go
panic("error message")
```

You can pass a string or any value to `panic()`.

---

## ✅ Example 1: Simple Panic

```go
package main

import "fmt"

func main() {
    fmt.Println("Before panic")
    panic("Something went wrong!")
    fmt.Println("After panic") // never runs
}
```

### 🧪 Output

```
Before panic
panic: Something went wrong!

goroutine 1 [running]:
main.main()
    ...stack trace...
```

---

## ✅ Example 2: Panic with Deferred Cleanup

```go
package main

import "fmt"

func main() {
    defer fmt.Println("Cleanup before crash")
    fmt.Println("Start")
    panic("Critical failure")
}
```

### 🧪 Output

```
Start
Cleanup before crash
panic: Critical failure
```

> 🔄 Deferred functions still run even when panic occurs.

---

## ✅ Example 3: Panic in a Function Call Chain

```go
package main

import "fmt"

func a() {
    b()
}

func b() {
    c()
}

func c() {
    panic("c panicked")
}

func main() {
    a()
}
```

### 🧪 Output

```
panic: c panicked

goroutine 1 [running]:
main.c()
...
main.main()
...
```

> 📉 Shows full stack trace leading to the panic.

---

## 🧷 Key Points Recap

| Concept         | Description                              |
| --------------- | ---------------------------------------- |
| `panic`         | Halts current execution                  |
| Stack unwinding | Deferred functions run in reverse order  |
| Stack trace     | Panic shows detailed function call stack |
| Use case        | Unexpected or unrecoverable errors       |

---

## 🧭 Summary

* `panic` is used to indicate **serious, unrecoverable errors**.
* It causes a **runtime crash**, but still **runs deferred functions**.
* It prints a **stack trace** to help debug the issue.
* Prefer normal error handling (`error` values) for expected failures.

> 🛠️ Use `panic` only when the program is in a state where continuing would be dangerous or meaningless.
