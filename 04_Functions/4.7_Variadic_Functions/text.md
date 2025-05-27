# 5.4. Variadic Functions in Go

---

## 📌 What are Variadic Functions?

A **variadic function** in Go is a function that can accept **a variable number of arguments** of the same type.

This is useful when you don’t know in advance how many arguments a function might receive.

---

## ❓ Why Use Variadic Functions?

* ✅ Useful for operations like `sum`, `print`, `log` where the number of inputs varies.
* ✅ Makes the function more flexible and easier to use.
* ✅ Helps in writing generic utilities.

> 🔄 Functions like `fmt.Println()` and `append()` are variadic.

---

## 🧾 Syntax of Variadic Function

```go
func functionName(args ...Type) {
    // args is a slice of Type
}
```

Inside the function, `args` behaves like a slice.

---

## ✅ Example 1: Sum Function

```go
package main

import "fmt"

func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    fmt.Println("Sum:", sum(1, 2, 3, 4))
    fmt.Println("Sum:", sum()) // valid, returns 0
}
```

### 🧪 Output

```
Sum: 10
Sum: 0
```

---

## ✅ Example 2: Mixing Fixed and Variadic Parameters

```go
package main

import "fmt"

func greet(greeting string, names ...string) {
    for _, name := range names {
        fmt.Println(greeting, name)
    }
}

func main() {
    greet("Hello", "Alice", "Bob")
    greet("Hi")
}
```

### 🧪 Output

```
Hello Alice
Hello Bob
```

---

## ✅ Example 3: Passing a Slice to Variadic Function

```go
package main

import "fmt"

func printNumbers(nums ...int) {
    for _, num := range nums {
        fmt.Println(num)
    }
}

func main() {
    numbers := []int{10, 20, 30}
    printNumbers(numbers...) // use ... to expand the slice
}
```

### 🧪 Output

```
10
20
30
```

---

## 🧷 Key Points Recap

| Concept                 | Description                             |
| ----------------------- | --------------------------------------- |
| Variadic function       | Accepts variable number of arguments    |
| `...type` syntax        | Defines variadic parameter              |
| Behaves like a slice    | Inside function, treated as a slice     |
| Can pass slice with ... | Expands slice into individual arguments |
| Fixed + variadic combo  | Fixed arguments must come first         |

---

## 🧭 Summary

* Variadic functions are a powerful way to accept flexible inputs.
* Internally, the variadic parameter is treated as a slice.
* You can mix fixed and variadic parameters, but variadic must be last.
* Use `...` to pass a slice into a variadic function.

> 🛠️ Variadic functions simplify the interface for utility functions like summing numbers, printing, or formatting data.
