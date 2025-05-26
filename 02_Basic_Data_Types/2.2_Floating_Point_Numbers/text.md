# 📘 2.2 Floating Point Numbers in Go

## 🧠 What Are Floating Point Numbers?

Floating point numbers are numbers that have a decimal point. They are used when precision with fractions is needed, such as in scientific calculations, measurements, or financial computations.

## 🧮 Types of Floating Point Numbers

Go provides two types of floating point numbers:

| Type    | Description            | Size    | Precision           |
| ------- | ---------------------- | ------- | ------------------- |
| float32 | 32-bit float           | 4 bytes | \~6 decimal places  |
| float64 | 64-bit float (default) | 8 bytes | \~15 decimal places |

`float64` is the default type for floating point literals.

## ✍️ Syntax

```go
var price float32 = 19.99         // Declaring a float32
var pi float64 = 3.1415926535     // Declaring a float64
```

You can also use shorthand declaration:

```go
weight := 65.5 // Inferred as float64
```

## ✅ Example Code

```go
package main

import "fmt"

func main() {
    var length float32 = 12.5    // Length in meters
    var width float32 = 7.8      // Width in meters

    area := length * width       // Calculate area

    // Print the results
    fmt.Println("Length:", length)
    fmt.Println("Width:", width)
    fmt.Println("Area:", area)
}
```

## 🔎 Output

```
Length: 12.5
Width: 7.8
Area: 97.5
```

## 🧹 Summary

| Feature   | Description                     |
| --------- | ------------------------------- |
| Types     | float32 and float64             |
| Precision | float64 is more precise         |
| Use Cases | Suitable for measurements, math |
| Shorthand | := infers float64 by default    |
