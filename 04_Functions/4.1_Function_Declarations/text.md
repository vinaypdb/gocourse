# 🔁 Function Declarations in Go

## 📘 Explanation

**📌 What is a Function Declaration?**

A **function declaration** in Go defines a named block of code that performs a specific task. It tells the compiler:

- What the function is called (its name)

- What inputs (parameters) it takes

- What outputs (return values) it produces

- What code to execute when it is called

**❓ Why Use Functions?**

- **Reusability**: Write once, use many times.

- **Modularity**: Break large problems into smaller chunks.

- **Readability**: Makes code cleaner and more understandable.

- **Testability**: Easier to test individual units of logic.


## 🧱 Syntax

```go
func functionName(param1 type1, param2 type2, ...) returnType {
    // function body
    // your code here
}

```

## ✅ Example: Function That Adds Two Numbers

```go
package main

import "fmt"

// add is a function that takes two integers and returns their sum.
func add(a int, b int) int {
    return a + b // return the result of adding a and b
}

func main() {
    result := add(3, 5) // calling the add function
    fmt.Println("Sum is:", result) // Output: Sum is: 8
}

```

## 🧪 Sample Output

```go
Sum is: 8

```

## ✅ Example 2: Return Multiple Values

```go
package main

import "fmt"

// divideAndRemainder returns both quotient and remainder
func divideAndRemainder(a, b int) (int, int) {
    return a / b, a % b
}

func main() {
    q, r := divideAndRemainder(10, 3)
    fmt.Println("Quotient:", q)
    fmt.Println("Remainder:", r)
}

```
## 🧪 Sample Output

```go
Quotient: 3
Remainder: 1

```

## ✅ Example 3: Grouped Parameter and Return Types

```go
package main

import "fmt"

// swap swaps the order of two strings
func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println("Swapped:", a, b)
}

```
## 🧪 Sample Output

```go
Swapped: world hello
```
## ✅ Example 4: Named Return Values

```go
package main

import "fmt"

// split splits a number into two parts
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return // returns x and y
}

func main() {
    a, b := split(18)
    fmt.Println("Split:", a, b)
}

```
## 🧪 Sample Output

```go
Split: 8 10
```
## ✅ Example 5: Ignoring Values Using `_`

```go
package main

import "fmt"

// divideAndRemainder returns both quotient and remainder
func divideAndRemainder(a, b int) (int, int) {
    return a / b, a % b
}

func main() {
    _, r := divideAndRemainder(10, 3)
    fmt.Println("Remainder only:", r)
}
```
## 🧪 Sample Output

```go
Remainder only: 1

```

## 💡 Key Points

| Concept                | Description                             |
| ---------------------- | --------------------------------------- |
| `func`                 | Keyword to declare a function           |
| Parameters             | Input values with types                 |
| Return type(s)         | What the function returns               |
| Multiple return values | Supported                               |
| Named return values    | Can be returned without explicit values |
| Blank identifier (`_`) | Ignore unused return values             |
| `main()`               | Entry point of every Go program         |




## 🧭 Summary

- Go functions begin with the `func` keyword.

- You can pass parameters and receive results.

- Supports **multiple returns** and **named return values**.

- Functions help organize, reuse, and simplify code logic.



