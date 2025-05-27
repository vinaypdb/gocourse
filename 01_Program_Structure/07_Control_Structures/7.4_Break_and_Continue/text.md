# 🔁  break and continue in Go

## 📘 Explanation

**📌 What are break and continue?**

Go provides two special control flow statements for use inside loops:

- `break`: Immediately exits the current loop.

- `continue`: Skips the remaining statements in the current iteration and moves to the next loop cycle.

These are useful for controlling loop flow dynamically.

## 🧱 Syntax

```go
for condition {
    if breakCondition {
        break
    }
    if skipCondition {
        continue
    }
    // other logic
}

```

## ✅ Example Code 1. Using break

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        if i == 5 {
            break // Exit the loop when i == 5
        }
        fmt.Println("i =", i)
    }
}

```

## 🧪 Sample Output

```go
i = 1
i = 2
i = 3
i = 4

```

## ✅ Example Code 2. Using continue

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        if i == 3 {
            continue // Skip when i == 3
        }
        fmt.Println("i =", i)
    }
}

```

## 🧪 Sample Output

```go
i = 1
i = 2
i = 4
i = 5

```
## 💡 🌀 Loop Control Flow

| Keyword    | Action                                |
| ---------- | ------------------------------------- |
| `break`    | Stops the loop completely             |
| `continue` | Skips current iteration and continues |


## 🔁 With Nested Loops and Labels
Go allows labels with break and continue to specify which loop to affect:

```go
outer:
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        if j == 2 {
            break outer // breaks the outer loop
        }
        fmt.Printf("i=%d, j=%d\n", i, j)
    }
}

```
## ❌ Common Mistakes Beginners Make

**1. Using break thinking it exits only if block**

```go
for i := 0; i < 5; i++ {
    if i == 3 {
        break  // ❌ exits the loop, not just the if
    }
    fmt.Println(i)
}

```
✅ Remember: `break` exits the whole loop.

**2. Misusing continue leading to missed logic**

```go
for i := 1; i <= 5; i++ {
    if i%2 == 0 {
        continue
    }
    fmt.Println(i, "is odd")
}

```
✅ Always place important logic **after** the `continue` check.

## 💡 Key Points

| Point                           | Description                    |
| ------------------------------- | ------------------------------ |
| `break` exits the loop          | Even if it's in a nested block |
| `continue` skips one iteration  | Does not exit the loop         |
| Labels help with nested loops   | e.g., `break outer`            |
| Helps avoid deeply nested logic | Makes loops clearer            |


## 🧠 Summary

- Use `break` to stop a loop early.

- Use `continue` to skip an iteration.

- Use labels if needed for nested loops.

- Both are crucial for writing clean, controlled loop logic.

