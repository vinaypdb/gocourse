# 📘 2.3 Complex Numbers in Go

## 🧠 What Are Complex Numbers?

Complex numbers have a real and an imaginary part. Go provides built-in support for complex numbers through two types:

* `complex64`: composed of two `float32` values
* `complex128`: composed of two `float64` values

## ✍️ Syntax

You can create a complex number using:

```go
var c complex64 = complex(1.5, 2.5) // real=1.5, imag=2.5
```

Or shorthand:

```go
c := complex(1, 2) // Default is complex128
```

You can extract real and imaginary parts using:

```go
realPart := real(c)
imagPart := imag(c)
```

## ✅ Example Code

```go
package main

import (
    "fmt"
)

func main() {
    // Declare complex numbers
    var a complex64 = complex(2.0, 3.0)
    var b complex64 = complex(1.5, 0.5)

    // Perform addition
    result := a + b

    // Print details
    fmt.Println("First:", a)
    fmt.Println("Second:", b)
    fmt.Println("Sum:", result)

    // Access real and imaginary parts
    fmt.Println("Real part of result:", real(result))
    fmt.Println("Imaginary part of result:", imag(result))
}
```

## 🔎 Output

```
First: (2+3i)
Second: (1.5+0.5i)
Sum: (3.5+3.5i)
Real part of result: 3.5
Imaginary part of result: 3.5
```

## 🧹 Summary

| Feature             | Description                                  |
| ------------------- | -------------------------------------------- |
| `complex64/128`     | Built-in complex number types                |
| `complex()`         | Function to construct complex numbers        |
| `real()` / `imag()` | Functions to extract parts of a complex      |
| Arithmetic          | Support for +, -, \*, / with complex numbers |
