# 🔄 If Statement in Go

## 📘 Explanation

**🧠 Purpose**

The `if` statement in Go is used to execute code based on a condition. If the condition evaluates to `true`, the code block inside the if statement is executed. If it evaluates to `false`, the code block is skipped.

## 🧱 Syntax

```
if condition {
    // Code to execute if condition is true
}

```

## ✅ Example: Simple If Statement

```go
package main

import "fmt"

func main() {
    x := 10
    if x > 5 {
        fmt.Println("x is greater than 5")
    }
}

```
**Explanation:**
- In the above example, since `x` is greater than `5`, the code inside the `if` block will run, and the output will be:

## 🧪 Sample Output

```
x is greater than 5
```
## 🔄 If-Else Statement
You can also use an `else` block to provide an alternative code path if the condition is `false`.
```go
if condition {
    // Code if condition is true
} else {
    // Code if condition is false
}

```
## ✅ Example: Simple If Statement
```go
package main

import "fmt"

func main() {
    x := 3
    if x > 5 {
        fmt.Println("x is greater than 5")
    } else {
        fmt.Println("x is less than or equal to 5")
    }
}

```
**Explanation:**

- The program checks if x is greater than 15 first, which is false.

- It then checks if x is exactly 10, which is true, so the output will be:

## 🧪 Sample Output
```go
x is exactly 10
```

## 🔄 Shortened If Statement

Go allows you to declare and initialize variables within an `if` statement.
```go
if variable := expression; condition {
    // Code to execute if condition is true
    // You can use variable here
}

```
## ✅ Example: Simple If Statement

```go
package main

import "fmt"

func main() {
    if x := 10; x > 5 {
        fmt.Println("x is greater than 5:", x)
    }
}


```
**Explanation:**

- The variable x is declared and initialized within the if statement, and the condition x > 5 is checked.

- The output will be:

## 🧪 Sample Output
```go
x is greater than 5: 10

```


## 🧩 Summary

- `if` in Go allows conditional execution of code.

- It can be extended with `else if` and `else` blocks.

- Go supports variable declarations directly inside `if` statements for scoped variables.


