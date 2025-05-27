# 4.2. Recursion in Go

---

## 📌 What is Recursion?

**Recursion** is a programming technique where a function calls **itself** to solve smaller instances of a problem.

Each recursive call works on a simpler or smaller input, and recursion continues until it reaches a **base case**, which stops the recursion.

---

## ❓ Why Use Recursion?

* ✅ **Simplifies problems** that have a naturally recursive structure (e.g., factorial, Fibonacci, tree traversal).
* ✅ Helps **break down complex tasks** into smaller subproblems.
* ✅ Reduces the need for explicit loops or stack management in some cases.

> 🤩 Many mathematical and algorithmic problems can be elegantly solved using recursion.

---

## 🧾 Syntax of Recursive Function

```go
func recursiveFunction(params...) returnType {
    if baseCaseCondition {
        return baseResult
    }
    return recursiveFunction(smallerInput) // Recursive call
}
```

---

## ✅ Example 1: Factorial Using Recursion

The factorial of a number `n` is defined as:
`n! = n × (n-1) × (n-2) × ... × 1`
with the base case: `0! = 1`

```go
package main

import "fmt"

// factorial returns the factorial of n using recursion
func factorial(n int) int {
    if n == 0 {
        return 1 // base case
    }
    return n * factorial(n-1) // recursive case
}

func main() {
    fmt.Println("Factorial of 5 is:", factorial(5))
}
```

### 🧪 Output

```
Factorial of 5 is: 120
```

---

## ✅ Example 2: Fibonacci Using Recursion

The Fibonacci sequence is defined as:
`F(0) = 0, F(1) = 1`
`F(n) = F(n-1) + F(n-2)`

```go
package main

import "fmt"

// fibonacci returns the nth Fibonacci number using recursion
func fibonacci(n int) int {
    if n <= 1 {
        return n // base cases
    }
    return fibonacci(n-1) + fibonacci(n-2) // recursive case
}

func main() {
    fmt.Println("Fibonacci number at position 6 is:", fibonacci(6))
}
```

### 🧪 Output

```
Fibonacci number at position 6 is: 8
```

---

## ✅ Example 3: Printing Countdown Using Recursion

```go
package main

import "fmt"

// countdown prints numbers from n to 1 using recursion
func countdown(n int) {
    if n == 0 {
        fmt.Println("Liftoff!")
        return
    }
    fmt.Println(n)
    countdown(n - 1) // recursive call
}

func main() {
    countdown(5)
}
```

### 🧪 Output

```
5
4
3
2
1
Liftoff!
```

---

## ⚠️ Base Case is Essential

Without a base case, recursion will go **infinitely** and cause a **stack overflow**.

---

## 🥇 Key Points Recap

| Concept            | Description                                               |
| ------------------ | --------------------------------------------------------- |
| Recursive function | A function that calls itself                              |
| Base case          | Stops recursion to prevent infinite loop                  |
| Recursive case     | The function calls itself with a simpler input            |
| Stack behavior     | Each call is stored on the call stack                     |
| Performance        | Can be less efficient than iteration without optimization |

---

## 🧽 Summary

* Recursion solves problems by breaking them into smaller subproblems.
* Always include a **base case** to stop recursion.
* Use recursion when the problem has a naturally recursive structure.
* Not all problems are efficient with recursion (e.g., naive Fibonacci is exponential).
* Recursion uses the **call stack**; too many recursive calls can cause **stack overflow**.
