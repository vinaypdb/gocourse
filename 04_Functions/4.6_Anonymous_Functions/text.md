# 5.3. Anonymous Functions in Go

---

## 📌 What are Anonymous Functions?

An **anonymous function** is a function **without a name**.

In Go, anonymous functions are also known as **function literals**. These functions can be:

* Assigned to variables
* Passed as arguments
* Returned from other functions
* Invoked immediately (IIFE)

---

## ❓ Why Use Anonymous Functions?

* ✅ Useful for short, simple functionality without polluting global scope.
* ✅ Helps write inline logic for things like filters, sorting, callbacks, etc.
* ✅ Makes code more readable and concise.
* ✅ Can **capture and use variables** from the outer scope (closures).

---

## 🧾 Syntax of Anonymous Function

```go
func(parameters) returnType {
    // function body
}(arguments) // optional immediate invocation
```

---

## ✅ Example 1: Assign Anonymous Function to Variable

```go
package main

import "fmt"

func main() {
    greet := func(name string) string {
        return "Hello, " + name
    }
    fmt.Println(greet("Go"))
}
```

### 🧪 Output

```
Hello, Go
```

---

## ✅ Example 2: Immediately Invoked Function Expression (IIFE)

```go
package main

import "fmt"

func main() {
    result := func(a, b int) int {
        return a + b
    }(5, 3)
    fmt.Println("Sum:", result)
}
```

### 🧪 Output

```
Sum: 8
```

---

## ✅ Example 3: Anonymous Function as Argument

```go
package main

import "fmt"

func operate(x, y int, op func(int, int) int) int {
    return op(x, y)
}

func main() {
    result := operate(10, 4, func(a, b int) int {
        return a - b
    })
    fmt.Println("Difference:", result)
}
```

### 🧪 Output

```
Difference: 6
```

---

## ✅ Example 4: Closures with Anonymous Functions

```go
package main

import "fmt"

func main() {
    counter := 0
    increment := func() int {
        counter++
        return counter
    }
    fmt.Println(increment()) // 1
    fmt.Println(increment()) // 2
}
```

### 🧪 Output

```
1
2
```

---

## 🧷 Key Points Recap

| Concept              | Description                            |
| -------------------- | -------------------------------------- |
| Anonymous function   | Function without a name                |
| Assigned to variable | Can store in a variable and call later |
| IIFE                 | Immediately invoked at definition      |
| Closure              | Captures and uses outer variables      |
| Used as argument     | Can pass directly to another function  |

---

## 🧭 Summary

* Anonymous functions are unnamed functions that can be defined inline.
* Useful for short, temporary logic, especially when defining **callbacks** or **one-time operations**.
* Can form **closures** by capturing outer variables.
* Can be **immediately invoked** or **assigned** for later use.

> 🛠️ Anonymous functions add power and flexibility to Go's functional capabilities.
