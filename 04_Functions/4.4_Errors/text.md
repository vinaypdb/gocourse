# 4.3. Errors in Go

---

## 📌 What are Errors?

An **error** in Go represents something that went wrong while a program was running. It is not a panic or crash, but a **controlled failure** that allows the program to handle and recover from issues gracefully.

Go uses a **built-in `error` interface** to manage error handling.

```go
type error interface {
    Error() string
}
```

This means any type that has an `Error() string` method is considered an error.

---

## ❓ Why Use Errors?

* ✅ Helps prevent program crashes by handling unexpected situations.
* ✅ Encourages explicit, clear handling of failures.
* ✅ Makes the program more reliable and robust.

> 🧩 Go favors explicit error handling over exceptions for simplicity and readability.

---

## 🧾 Syntax: Returning and Checking Errors

```go
func someFunction() (resultType, error) {
    if somethingWentWrong {
        return defaultResult, errors.New("error message")
    }
    return result, nil
}
```

To check the error:

```go
result, err := someFunction()
if err != nil {
    // handle error
    fmt.Println("Error:", err)
} else {
    // use result
}
```

> 🔍 The convention in Go is to return the error as the **last return value**.

---

## ✅ Example 1: Basic Error Handling

```go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

### 🧪 Output

```
Error: cannot divide by zero
```

---

## ✅ Example 2: Custom Error Type

You can define your own custom error types by implementing the `Error()` method.

```go
package main

import "fmt"

type MyError struct {
    Code int
    Message string
}

func (e MyError) Error() string {
    return fmt.Sprintf("[Error %d]: %s", e.Code, e.Message)
}

func riskyOperation(fail bool) error {
    if fail {
        return MyError{Code: 400, Message: "Operation failed"}
    }
    return nil
}

func main() {
    err := riskyOperation(true)
    if err != nil {
        fmt.Println("Caught error:", err)
    }
}
```

### 🧪 Output

```
Caught error: [Error 400]: Operation failed
```

---

## 🧷 Key Points Recap

| Concept              | Description                        |
| -------------------- | ---------------------------------- |
| `error` interface    | Standard way to represent errors   |
| `errors.New()`       | Creates a new error with a message |
| `nil`                | Means no error occurred            |
| Custom error type    | Any struct implementing `Error()`  |
| Error as last return | Go convention for returning errors |

---

## 🧭 Summary

* Go uses the `error` interface for explicit, clear error handling.
* Always check if `err != nil` before using the result.
* Errors are returned, not thrown.
* You can define **custom error types** for more detailed error reporting.
* Follow Go conventions: errors should be the **last** value returned.

> 🛠️ Good error handling makes your Go programs reliable and production-ready.
