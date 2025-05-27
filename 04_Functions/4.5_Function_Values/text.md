# 5.2. Function Values in Go

---

## 📌 What are Function Values?

In Go, **functions are first-class values**. This means:

* You can **assign a function to a variable**.
* You can **pass a function as an argument** to another function.
* You can **return a function** from a function.

These are called **function values** or **function literals**.

---

## ❓ Why Use Function Values?

* ✅ Enables writing **higher-order functions**.
* ✅ Useful for callbacks, filters, and flexible behavior.
* ✅ Helps create **modular and reusable code**.
* ✅ Encourages concise inline logic with **anonymous functions**.

---

## 🧾 Syntax: Function Value Assignment

```go
func add(a, b int) int {
    return a + b
}

var operation func(int, int) int = add
```

Or using type inference:

```go
operation := add
```

---

## ✅ Example 1: Assigning a Function to a Variable

```go
package main

import "fmt"

func multiply(x, y int) int {
    return x * y
}

func main() {
    operation := multiply
    fmt.Println("Result:", operation(4, 5)) // 4 * 5 = 20
}
```

### 🧪 Output

```
Result: 20
```

---

## ✅ Example 2: Anonymous Function (Function Literal)

```go
package main

import "fmt"

func main() {
    square := func(n int) int {
        return n * n
    }
    fmt.Println("Square of 6:", square(6))
}
```

### 🧪 Output

```
Square of 6: 36
```

---

## ✅ Example 3: Passing a Function as Argument

```go
package main

import "fmt"

func apply(x, y int, op func(int, int) int) int {
    return op(x, y)
}

func add(a, b int) int {
    return a + b
}

func main() {
    result := apply(10, 5, add)
    fmt.Println("Sum:", result)
}
```

### 🧪 Output

```
Sum: 15
```

---

## ✅ Example 4: Returning a Function

```go
package main

import "fmt"

func makeMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func main() {
    times3 := makeMultiplier(3)
    fmt.Println("3 times 7:", times3(7))
}
```

### 🧪 Output

```
3 times 7: 21
```

---

## 🧷 Key Points Recap

| Concept               | Description                                       |
| --------------------- | ------------------------------------------------- |
| Function value        | A function assigned to a variable                 |
| Anonymous function    | Function without a name, often used inline        |
| Higher-order function | A function that takes or returns another function |
| Closures              | Returned function captures surrounding variables  |
| Function type         | Specified using `func(paramTypes) returnType`     |

---

## 🧭 Summary

* Functions in Go can be assigned to variables, passed, and returned.
* Function values are powerful for abstraction and modular design.
* Use **anonymous functions** for concise logic.
* Use **closures** to retain state between function calls.

> 🛠️ Function values are the foundation of flexible and expressive Go programs.
