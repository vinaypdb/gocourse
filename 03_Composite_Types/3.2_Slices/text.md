# 📘 3.2 Slices in Go

🧠 **What Are Constants?**

A **slice** is like a flexible, dynamic view (or window) into an underlying array. Unlike arrays, slices can **grow or shrink** in size as needed.

You can think of a slice as a reference to a part (or all) of an array. It keeps track of:

- The starting point (where the slice begins),

- The length (how many elements it currently holds),

- The capacity (maximum size it can grow to without reallocating).

**Why Use Slices Instead of Arrays?**

- Arrays are fixed in size — you must specify their size upfront.

- Slices can grow and shrink, making them more flexible and easier to work with.

- Slices are much more common in Go programs because of this flexibility.


## ✍️ Syntax

```go
var sliceName []Type
```

- Note there is no size inside the brackets [].

- The slice starts empty until you assign or append elements.


## ✅ Example Code

```go
package main

import "fmt"

func main() {
    // Declare a slice of integers (initially nil, length 0)
    var numbers []int

    // Check if slice is nil or empty
    fmt.Println("Is 'numbers' slice nil?", numbers == nil)
    fmt.Println("Length:", len(numbers), "Capacity:", cap(numbers))

    // Append elements to the slice (slice grows automatically)
    numbers = append(numbers, 10)
    numbers = append(numbers, 20, 30, 40)

    // Print the slice and its length and capacity
    fmt.Println("Numbers slice:", numbers)
    fmt.Println("Length:", len(numbers), "Capacity:", cap(numbers))

    // Create a slice from an array
    arr := [5]int{1, 2, 3, 4, 5}
    sliceFromArr := arr[1:4] // slice contains elements arr[1], arr[2], arr[3]
    fmt.Println("Slice from array:", sliceFromArr)

    // Modify a slice element — this also changes the underlying array
    sliceFromArr[0] = 100
    fmt.Println("Modified slice:", sliceFromArr)
    fmt.Println("Underlying array after modification:", arr)
}


```

## 🔎 Output

```
Is 'numbers' slice nil? true
Length: 0 Capacity: 0
Numbers slice: [10 20 30 40]
Length: 4 Capacity: 4
Slice from array: [2 3 4]
Modified slice: [100 3 4]
Underlying array after modification: [1 100 3 4 5]

```

## 🔐 Explanation of the Example

- Declared a slice named `numbers`. Initially, it’s `nil` (no underlying array).

Used `append()` to add elements. The slice grows dynamically.

Printed length and capacity using `len()` and `cap()`.

Created a slice from an existing array (`arr[1:4]` means slice elements from index 1 up to but not including 4).

Changing the slice element also changes the corresponding element in the underlying array because the slice refers to the same data.


## Key Points to Remember

- Slices are more flexible than arrays: they can grow and shrink.

- Slices have **length and capacity;** capacity is the max size before needing to allocate a new underlying array.

- Use `append(slice, elements...)` to add elements.

- Slices are reference types — copying a slice copies the reference, not the underlying data.

- Modifying a slice modifies the underlying array.

- You can create slices from arrays or other slices using `slice[start:end]` syntax.

- A nil slice has length and capacity 0 and no underlying array.



