# 🔁 Traditional for Loop in Go

## 📘 Explanation

**🧠 Purpose**

A **traditional** `for` **loop** is used when you want to repeat a block of code a fixed number of times. It's similar to the classic `for` loop syntax seen in languages like C, C++, or Java.

## 🤔 Why Use a Traditional for Loop?

Use it when:

- You know **exactly how many times** you want the loop to run.

- You need to control the index (loop counter).

- You want to run a block of code a set number of times with clear initialization, condition checking, and updating logic.

## 🧱 Syntax

```
for initialization; condition; post {
    // loop body
}

```
**Components:**

- **Initialization:** Runs once at the beginning (e.g., `i := 0`)

- **Condition:** Checked before each iteration (loop continues while `condition` is true)

- **Post:** Executes after each iteration (e.g., `i++`)

## ✅ Example Code 

```go
package main

import "fmt"

func main() {
    // Print numbers from 1 to 5
    for i := 1; i <= 5; i++ {
        fmt.Println("Iteration:", i)
    }
}

```

## 🧪 Sample Output

```
Iteration: 1
Iteration: 2
Iteration: 3
Iteration: 4
Iteration: 5

```
**Explanation:**

- `i := 1` → initializes loop variable `i` to 1.

- `i <= 5` → condition: continue looping while `i` is ≤ 5.

- `i++` → increment `i` after each iteration.

## 🔄 Common Mistakes Beginners Make

**1.Missing semicolons (;)**
```
go
for i := 0 i < 5 i++ {  // ❌ Invalid syntax (missing `;`)

```
**✅ Correct:**

```
for i := 0; i < 5; i++ {

```
**2.Forgetting to increment (`i++`)**
```go
for i := 0; i < 5; {
    fmt.Println(i)
    // ❌ Infinite loop — i is never incremented!
}

```
**✅ Correct:**

```
for i := 0; i < 5; i++ {

```
**3.Incorrect loop condition**
```go
for i := 1; i >= 5; i++ {
    // ❌ Will never execute — condition is initially false
}

```
**✅ Correct:**

```
for i := 1; i <= 5; i++ {

```
**4.Using `=` instead of `==` in condition**
```go
for i := 0; i = 5; i++ {  // ❌ Assignment, not comparison

```
**✅ Correct:**

```
for i := 0; i == 5; i++ {  // ✅ Comparison

```

## 💡 Key Points
| Point                                         | Description                     |
| --------------------------------------------- | ------------------------------- |
| `for` is the **only** loop keyword in Go      | No `while` or `do-while`        |
| All 3 parts (`init; cond; post`) are optional | Allows flexible loop forms      |
| Curly braces `{}` are **required**            | Even for single-line loops      |
| Loop variable is block-scoped                 | Only accessible inside the loop |
| Can be combined with `break` or `continue`    | For advanced control flow       |



## 🧠 Summary

- A **traditional** for **loop** is perfect when you know how many times to iterate.

- It follows the format: `for init; condition; post {}`.

- Avoid beginner pitfalls like missing semicolons or incorrect conditions.

- Go doesn’t have `while`, but this loop can serve the same purpose with adjustments.