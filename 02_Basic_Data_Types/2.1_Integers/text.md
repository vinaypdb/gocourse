# 📘 2.1 Integers in Go

## 🧠 What Are Integers?

In Go, integers are whole numbers without any decimal points. They're used for counting, indexing, looping, and more.

## 🧮 Types of Integer Data Types

Go supports different integer types based on size and whether they are signed or unsigned:

| Type   | Description                 | Size                            |
| ------ | --------------------------- | ------------------------------- |
| int    | Signed integer (platform)   | 32 or 64 bits                   |
| int8   | Signed 8-bit integer        | -128 to 127                     |
| int16  | Signed 16-bit integer       | -32,768 to 32,767               |
| int32  | Signed 32-bit integer       | -2,147,483,648 to 2,147,483,647 |
| int64  | Signed 64-bit integer       | Big range                       |
| uint   | Unsigned integer (platform) | 32 or 64 bits                   |
| uint8  | Unsigned 8-bit integer      | 0 to 255 (aka byte)             |
| uint16 | Unsigned 16-bit integer     | 0 to 65,535                     |
| uint32 | Unsigned 32-bit integer     | 0 to 4,294,967,295              |
| uint64 | Unsigned 64-bit integer     | Big range                       |

## ✍️ Syntax

```go
var age int = 25              // Declaring an int with explicit type
var smallNumber int8 = -100   // Declaring an int8 with a small negative number
var bigNumber uint64 = 100000000000 // Declaring a uint64 with a large positive number
```

You can also use shorthand declaration:

```go
count := 10 // Declares and initializes an int with value 10
```

## ✅ Example Code

```go
package main

import "fmt"

func main() {
    var apples int = 10              // Number of apples (signed integer)
    var oranges uint = 5             // Number of oranges (unsigned integer)

    // Convert oranges (uint) to int for addition
    total := apples + int(oranges)  // Type conversion from uint to int before addition

    // Print the results
    fmt.Println("Apples:", apples)        // Output: Apples: 10
    fmt.Println("Oranges:", oranges)      // Output: Oranges: 5
    fmt.Println("Total Fruits:", total)   // Output: Total Fruits: 15
}
```

## 🔎 Output

```
Apples: 10
Oranges: 5
Total Fruits: 15
```

## 🧹 Summary

| Feature       | Description                        |
| ------------- | ---------------------------------- |
| Integer Types | Various signed and unsigned types  |
| Size Options  | From 8-bit to 64-bit               |
| Type Casting  | Needed when mixing signed/unsigned |
| Shorthand     | := for quick declarations          |
