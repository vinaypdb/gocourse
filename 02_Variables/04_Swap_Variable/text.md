# 🔄 Swaping variables in Go

## 📘 Explanation

In Go, you can swap the values of two variables without using a temporary variable. This is possible because Go supports multiple assignment, allowing two variables to exchange values in a single line.

## ✍️ Syntax

```go
a, b = b, a
```
This statement does two things at once:

- The value of `b` is assigned to `a`

- The value of `a` is assigned to `b`

## ✅ Example Code

```go
package main

import "fmt"

func main() {
    a := 10
    b := 20

    fmt.Println("Before swap:")
    fmt.Println("a:", a)
    fmt.Println("b:", b)

    // Swap values
    a, b = b, a

    fmt.Println("After swap:")
    fmt.Println("a:", a)
    fmt.Println("b:", b)
}


```

## 🧪 Sample Output

```
Before swap:
a: 10
b: 20
After swap:
a: 20
b: 10

```

## 🧩 Summary

| Feature             | Description                                |
|---------------------|--------------------------------------------|
| Multiple Assignment | Allows assigning multiple variables at once        |
| No Temp Variable             | 	No need for a third variable to swap values                        |
| Clean Syntax	     | Code is concise and easy to understand  |
