# 🔁 goto Statement in Go

## 📘 Explanation

**📌 What is goto ?**

The `goto` statement provides a way to **jump** to another line in the same function, identified by a **label**.

It is a **low-level control structure** that can simplify certain branching logic (e.g., breaking out of deeply nested loops), but it's generally used sparingly to avoid unreadable code.

**🤔  Why Use goto?**

- To **break out** of deeply nested loops or conditions.

- For **early exit** in some complex flow control.

- Not recommended for general-purpose code – can make logic hard to follow.

## ⚠️ Note

Use of `goto` is **discouraged** in high-level code due to poor readability. Prefer `break`, `continue`, and functions for clean control flow.



## 🧱 Syntax

```go
goto labelName

...

labelName:
    // statements

```

- Label must be in the same function.

- Label must be followed by a statement.

- No backward jumps to a point before a defer statement.

## ✅ Example: Switch with Expression

```go
package main

import "fmt"

func main() {
    fmt.Println("Start")

    goto end

    fmt.Println("This will be skipped")

end:
    fmt.Println("End")
}

```

## 🧪 Sample Output

```go
Start  
End
```

## ✅ Example: Breaking Nested Loops

```go
package main

import "fmt"

func main() {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i*j == 2 {
                goto exit
            }
            fmt.Println("i:", i, "j:", j)
        }
    }

exit:
    fmt.Println("Exited nested loops")
}

```


## 💡 Key Points

| Feature            | Description                           |
| ------------------ | ------------------------------------- |
| `goto`             | Jumps to a label in the same function |
| Use cases          | Breaking nested loops, rare logic     |
| Not used for logic | Avoid for general control flow        |
| Same function only | Cannot jump between functions         |



## 🧠 Summary

- `goto` provides a jump to a labeled statement.

- Use it **sparingly and carefully** to avoid unreadable code.

- Prefer structured control flows (loops, functions, conditionals) instead.

