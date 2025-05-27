# 🔁 🔁 While-Style for Loop in Go

## 📘 Explanation

**🧠 Purpose**

Go does not have a separate `while` keyword. Instead, a `while`-style loop is implemented using the `for` loop **without the initialization and post statements**.
It keeps running as long as the condition evaluates to true.

## 🤔 Why Use a While-Style Loop?

Use it when:

- The number of iterations is not fixed.

- The loop depends on a dynamic condition, such as waiting for user input, reaching the end of a file, or a specific value being met.

- You want the logic and update step to be more flexible inside the loop body.

## 🧱 Syntax

```go
for condition {
    // loop body
}
```
This is functionally equivalent to:
```go
while (condition) {
    // loop body
}

```
in C, Java, or Python.


## ✅ Example Code 

```go
package main

import "fmt"

func main() {
    i := 0
    // Loop while i < 5
    for i < 5 {
        fmt.Println("i =", i)
        i++ // Update step inside the loop
    }
}

```

## 🧪 Sample Output

```go
i = 0
i = 1
i = 2
i = 3
i = 4

```

## 🔄 Common Mistakes Beginners Make

**1.Forgetting to update the loop variable**
```
i := 0
for i < 5 {
    fmt.Println(i)
    // ❌ No increment — infinite loop
}

```
**✅ Correct:**

```
i := 0
for i < 5 {
    fmt.Println(i)
    i++  // ✅ Proper update
}

```
**2.Incorrect or never-true condition**
```go
i := 10
for i < 5 {
    fmt.Println(i)
    i++
}
// ❌ Loop body never runs — condition is false initially

```
**✅ Correct:**

Ensure the initial value of `i` allows the condition to be true at some point.

**3.Modifying condition variable incorrectly**
```go
i := 0
for i < 5 {
    i = i + 2 // Jumps over some values
    fmt.Println(i)
}
// ✅ But might not behave as expected if increment is wrong

```


## 💡 Key Points

| Point                                     | Description                             |
| ----------------------------------------- | --------------------------------------- |
| Go has no `while` loop                    | Use `for condition` instead             |
| Initialization and update done separately | More control inside loop body           |
| Break and continue are still valid        | Use `break` to exit, `continue` to skip |
| Risk of infinite loops                    | If condition never becomes false        |



## 🧠 Summary

- A while-style loop uses `for` with only a condition.

- Ideal when the number of iterations is unknown or based on runtime conditions.

- Always make sure the condition eventually becomes `false` to avoid infinite loops.

- Keep your loop variables and update logic clear and well-placed.