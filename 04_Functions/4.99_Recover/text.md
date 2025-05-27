# Recover in Go

---

## 📌 What is recover?

`recover` is a built-in Go function that allows a program to regain control of a panicking goroutine. When a Go program encounters a **panic** (an unexpected runtime error), it normally stops execution and begins unwinding the stack, running deferred functions on the way.

`recover` is used inside a deferred function to **catch the panic and stop the program from crashing**, allowing the program to continue executing or handle the error gracefully.

---

## ❓ Why Use recover?

- To handle unexpected errors without crashing the entire program.

- To clean up resources or save state before exiting or resuming execution.

- To implement robust libraries or servers that do not stop because of a single goroutine panic.

- To perform custom error handling or logging during panics.

**Important:** `recover` only works if called inside a deferred function; otherwise, it returns `nil`.

---

## 🧾 Syntax 

```go
func recover() interface{}
```
---

## How `recover` Works: Basic Flow
  1. A function panics (panic(value)).

  2. The runtime starts unwinding the stack, running deferred functions.

  3. If a deferred function calls recover() during this unwinding, it receives the panic     value and stops the panic.

  4. The program can then continue normally (or handle the error accordingly).

## ✅ Example Code

```go
package main

import "fmt"

// function that causes a panic
func riskyFunction() {
    fmt.Println("Starting risky function")
    panic("something went wrong!") // triggers a panic
    fmt.Println("This line will never run")
}

// function with deferred recover to handle panic
func safeFunction() {
    defer func() {
        if r := recover(); r != nil {
            // recover returns the panic value
            fmt.Println("Recovered from panic:", r)
        }
    }()
    
    fmt.Println("Calling riskyFunction safely")
    riskyFunction() // panic will happen here, but recover will catch it
    fmt.Println("After riskyFunction call") // this will not run if panic not recovered
}

func main() {
    fmt.Println("Program started")
    safeFunction()
    fmt.Println("Program continues after panic recovery")
}

```

### 🧪 Output

```
Program started
Calling riskyFunction safely
Starting risky function
Recovered from panic: something went wrong!
Program continues after panic recovery

```

---

**Explanation of the Example**

- `riskyFunction()` panics.

- `safeFunction()` defers an anonymous function with a call to `recover()`.

- When `riskyFunction()` panics, Go starts unwinding the stack and calls the deferred function.

- Inside the deferred function, `recover()` captures the panic value (`"something went wrong!"`).

- Because the panic was recovered, the program does not crash and continues execution.

- The program prints the recovered message and continues after `safeFunction()` returns.


## Key Points

- `recover` only works inside deferred functions. Calling it anywhere else returns `nil`.

- `recover` captures the panic value and stops the unwinding process.

- If you don't recover a panic, the program will crash.

- Use `recover` **sparingly and carefully** — catching every panic can hide bugs.

- It's common to use `recover` to make **servers or libraries resilient** by handling panics internally.

- The combination of `panic` and `recover` is Go’s way of handling **exceptions** but in a controlled manner.







