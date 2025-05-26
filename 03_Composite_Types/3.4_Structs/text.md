# 📘 3.4 Structs in Go

🧠 **What is a Struct?**

A struct is a composite data type in Go that groups **multiple fields** together.
Each field can have a different type — making structs perfect for modeling real-world objects.

Think of a struct as a **template** for a custom data type.

For example: A `Person` can have a `name`, `age`, and `email`. These can all be bundled together in a struct.

## ✍️ Syntax

```go
type StructName struct {
    field1 fieldType1
    field2 fieldType2
    ...
}

```
- `type` — declares a new named type.

- `StructName` — the name of the struct.

- Fields are declared inside the `{}` block.

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

// Declare a struct named Person
type Person struct {
    name  string
    age   int
    email string
}

func main() {
    // Create a struct instance and assign values
    var p1 Person
    p1.name = "Alice"
    p1.age = 28
    p1.email = "alice@example.com"

    // Print the struct
    fmt.Println("Person 1:", p1)

    // Access individual fields
    fmt.Println("Name:", p1.name)
    fmt.Println("Age:", p1.age)

    // Create a struct using a literal
    p2 := Person{name: "Bob", age: 32, email: "bob@example.com"}
    fmt.Println("Person 2:", p2)

    // You can also use positional field values (not recommended for clarity)
    p3 := Person{"Charlie", 40, "charlie@example.com"}
    fmt.Println("Person 3:", p3)
}

```

## 🔎 Output

```
Person 1: {Alice 28 alice@example.com}
Name: Alice
Age: 28
Person 2: {Bob 32 bob@example.com}
Person 3: {Charlie 40 charlie@example.com}

```

## 🔐 Explanation of the Example

- Defined a `Person` struct with 3 fields: `name`, `age`, `email`.

- Created struct instances (`p1`, `p2`, `p3`) using different styles.

- Accessed and printed struct fields with dot notation (e.g., `p1.name`).




## Key Points to Remember

- A struct is a way to **group related values** of different types.

- Structs can be **declared with field names and types.**

- Use dot (`.`) notation to access or assign individual fields.

- You can create structs using:

    - Field-by-field assignment

    - Struct literals with named fields

    - Struct literals with positional values

- Structs are **value types** — assigning a struct to another copies all its fields.



