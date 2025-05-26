# 📘 3.3 Maps in Go

🧠 **What is a Map?**

A **map** is a built-in data type in Go that stores **key-value pairs**.
It’s like a dictionary or hash table where you look up values using unique keys.

For example:

- Keys can be strings, integers, or any comparable type.

- Values can be of any type.

Maps let you quickly find a value based on its key.

## ✍️ Syntax

```go
var mapName map[keyType]valueType

```
- `mapName` — the name of the map variable.

- `keyType` — the type of the keys (e.g., `string`, `int`).

- `valueType` — the type of the values stored.

**How to Create a Map**

- Using `make()` function (most common way):

```go
mapName := make(map[keyType]valueType)

```
- Or with a map literal to initialize with values:

```go
mapName := map[keyType]valueType{
    key1: value1,
    key2: value2,
}


```


## ✅ Example Code

```go
package main

import "fmt"

func main() {
    // Create a map with string keys and int values
    ages := make(map[string]int)

    // Add key-value pairs
    ages["Alice"] = 25
    ages["Bob"] = 30
    ages["Charlie"] = 35

    // Print the map
    fmt.Println("Ages map:", ages)

    // Access a value by key
    fmt.Println("Age of Bob:", ages["Bob"])

    // Check if a key exists
    age, exists := ages["Diana"]
    if exists {
        fmt.Println("Age of Diana:", age)
    } else {
        fmt.Println("Diana is not found in the map.")
    }

    // Delete a key-value pair
    delete(ages, "Alice")
    fmt.Println("After deleting Alice:", ages)
}

```

## 🔎 Output

```
Ages map: map[Alice:25 Bob:30 Charlie:35]
Age of Bob: 30
Diana is not found in the map.
After deleting Alice: map[Bob:30 Charlie:35]

```

## 🔐 Explanation of the Example

- Created a map named `ages` with `string` keys and `int` values.

- Added some entries (key-value pairs).

- Printed the whole map.

- Accessed the value for the key `"Bob"`.

- Checked if `"Diana"` exists in the map using the comma-ok idiom.

- Deleted `"Alice"` from the map using the built-in `delete()` function.




## Key Points to Remember

- Maps store **key-value pairs**.

- Keys must be **comparable** types (like strings, ints, booleans).

- Maps are **reference types** (changes affect the original map).

- Use `make()` to create a map before adding entries.

- To check if a key exists, use the **comma-ok idiom:** `value, ok := map[key]`.

- Use `delete(map, key)` to remove an entry.

- The order of iteration over a map is `random` — it does not preserve insertion order.



