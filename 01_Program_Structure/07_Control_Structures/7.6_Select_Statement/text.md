# 🔁 select Statement in Go (Concurrency)

## 📘 Explanation

**📌 What is select?**

The `select` statement is used with **channels** to **wait on multiple communication operations**. It allows a Goroutine to wait until **one of several channel operations can proceed**.

**🤔 Why Use select?**

- Handle **multiple channels** concurrently.

- Build **non-blocking, timeout-aware**, and **responsive** programs.

- Avoid complex synchronization using manual logic.


## 🧱 Syntax

```go
select {
case val := <-ch1:
    // receive from ch1
case ch2 <- val:
    // send val to ch2
default:
    // run if no channel is ready
}

```

## ✅ Example 1: Multiple Receives

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    // Simulate goroutine sending to ch1 after 2 seconds
    go func() {
        time.Sleep(2 * time.Second)
        ch1 <- "from ch1"
    }()

    // Simulate goroutine sending to ch2 after 1 second
    go func() {
        time.Sleep(1 * time.Second)
        ch2 <- "from ch2"
    }()

    // Wait until one of the channels is ready
    select {
    case msg1 := <-ch1:
        fmt.Println("Received:", msg1)
    case msg2 := <-ch2:
        fmt.Println("Received:", msg2) // This will print first
    }
}


```

## 🧪 Sample Output

```go
Received: from ch2
```

## ✅ Example 2: Non-blocking Channel with default

```go
package main

import "fmt"

func main() {
    ch := make(chan string)

    select {
    case msg := <-ch:
        fmt.Println("Received:", msg)
    default:
        // This executes because channel is empty and non-blocking
        fmt.Println("No channel ready, moving on")
    }
}

```
## ✅ Example 3: Timeout with time.After

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string)

    select {
    case msg := <-ch:
        fmt.Println("Received:", msg)
    case <-time.After(2 * time.Second):
        // Timeout case executes if no message is received in 2 seconds
        fmt.Println("Timeout!")
    }
}

```

## 💡 Key Points

| Concept        | Description                           |
| -------------- | ------------------------------------- |
| `select`       | Multiplexes communication on channels |
| `default` case | Makes `select` non-blocking           |
| `time.After`   | Allows timeout handling               |
| Case selection | Random if multiple channels are ready |
| Channel safety | Check for closure using `ok` syntax   |



## 🧠 Summary

- `select` helps manage **multiple channels concurrently**.

- It blocks until a **case is ready**, unless `default` is present.

- Ideal for **timeouts, channel multiplexing**, and **responsive** apps.

- Always ensure **channels are properly closed and checked**.



