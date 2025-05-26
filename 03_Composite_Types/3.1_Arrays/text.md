# 📘 3.1 Arrays in Go

🧠 **What Are Constants?**

An **array** is a fixed-size collection of elements of the **same type** stored in contiguous memory. Think of it like a row of boxes where each box holds a value, and all boxes are of the same kind (e.g., all integers or all strings).

Once you create an array, its size cannot change — it always holds a fixed number of elements.


## ✍️ Syntax

```go
var arrayName [size]Type
```

- `var` — keyword to declare a variable.

- `arrayName` — name of the array variable.

- `[size]` — the number of elements the array can hold (fixed size).

- `Type` — the type of elements the array will store (e.g., `int`, `string`, `float64`).


## ✅ Example Code

```go
package main

import "fmt"

func main() {
    // Declare an array of 5 integers
    var numbers [5]int

    // Assign values to each element of the array using index (index starts at 0)
    numbers[0] = 10
    numbers[1] = 20
    numbers[2] = 30
    numbers[3] = 40
    numbers[4] = 50

    // Print the whole array
    fmt.Println("Numbers array:", numbers)

    // Access and print a single element
    fmt.Println("Element at index 2:", numbers[2])

    // Array length
    fmt.Println("Length of the array:", len(numbers))

    // Declare and initialize an array in one line
    fruits := [3]string{"apple", "banana", "cherry"}
    fmt.Println("Fruits array:", fruits)
}

```

## 🔎 Output

```
Days in a week: 7
Welcome to Go!

```

## 🔐 Explanation of the Example

- We declared an integer array named numbers that can hold 5 elements.

- We assigned values to each element by referring to its index (from 0 to 4).

- Printed the entire array using `fmt.Println`.

- Printed a specific element (`numbers[2]`).

- Used the built-in function `len()` to get the length of the array.

- Also showed how to declare and initialize an array in one line with fruit names.


## Key Points to Remember

- Arrays have **fixed size**. The size is part of the type (e.g., `[5]int` is different from `[3]int`).

- Array elements are accessed using **zero-based indices** (first element is at index 0).

- Arrays store elements of the **same type only**.

- Arrays in Go are **value types** — when you assign one array to another, it copies all elements.

- Use `len(array)` to get the number of elements.

- For flexible, resizable collections, consider using **slices** (more advanced topic).




