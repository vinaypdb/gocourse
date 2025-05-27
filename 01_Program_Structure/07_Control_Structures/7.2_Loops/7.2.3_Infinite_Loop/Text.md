# 🔁 Infinite Loop in Go

## 📘 Explanation

**🧠 Purpose**

An **infinite loop** is a loop that **never ends on its own** because its condition is always `true`. In Go, this is written using a `for` statement with **no condition** at all:

**🤔 Why Use an Infinite Loop?**

Use it when:

- You're waiting for a signal, event, or input (like a server).

- You’re building a **continuous service** (e.g., network listener, background task).

- You want a loop that exits only when a specific internal condition is met.

## 🧱 Syntax

```go
for {
    // loop body
}

```

## ✅ Example Code 

```go
package main

import "fmt"

func main() {
    count := 0

    for {
        fmt.Println("Count:", count)
        count++

        // Manually break the loop after 3 iterations
        if count == 3 {
            fmt.Println("Breaking the loop")
            break
        }
    }
}


```

## 🧪 Sample Output

```go
Count: 0
Count: 1
Count: 2
Breaking the loop

```
## 🔄 Real-World Use Case: Server Loop
```go
for {
    conn, err := listener.Accept()
    if err != nil {
        log.Println("Connection error:", err)
        continue
    }
    go handleConnection(conn)
}


```
This runs a server loop that never ends unless the program is terminated.



## 🔄 Common Mistakes Beginners Make

**1. No exit condition — accidental infinite loop**
```go
for {
    fmt.Println("Hello")
}
// ❌ This never stops unless manually killed (Ctrl+C)

```

**✅ Fix:** Always ensure a condition to exit (like break).

**2. Wrong use of `break` or forgetting it**

```go
i := 0
for {
    i++
    // ❌ Forgetting break leads to infinite loop
}

```
✅ **Fix:**

```go
i := 0
for {
    if i >= 5 {
        break
    }
    fmt.Println(i)
    i++
}

```


## 💡 Key Points

| Point                                              | Description                        |
| -------------------------------------------------- | ---------------------------------- |
| `for {}` is a valid infinite loop                  | No condition means it runs forever |
| Must use `break`, `return`, or `os.Exit()` to stop | Or it will loop infinitely         |
| Common in server applications or background tasks  | Especially in goroutines           |
| Dangerous if not handled correctly                 | May freeze your app or consume CPU |



## 🧠 Summary

- An infinite loop is written as `for {}` in Go.

- It is powerful but must be used carefully.

- Always include a way to exit or control it.

- Ideal for tasks that need to run continuously until explicitly stopped.

