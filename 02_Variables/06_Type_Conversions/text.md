# 🔄 6. Type Conversion in Go

## 📘 Explanation

In Go, type conversion (also called type casting) is used to explicitly convert a value from one data type to another. Go does not support implicit type conversion, meaning you must manually convert between different types.

This is important when performing operations between variables of different types (e.g., int and float64), or when formatting data.

## ✍️ Syntax

```go
convertedValue := TypeName(originalValue)
```

## ✍️ For example:

```go
var x int = 10
var y float64 = float64(x)
```

## ✅ Example

```go
package main

import "fmt"

func main() {
    var a int = 42
    var b float64 = float64(a)             // int to float64
    var c string = fmt.Sprintf("%d", a)    // int to string

    fmt.Println("Integer:", a)
    fmt.Println("Float:", b)
    fmt.Println("String:", c)
}
```

## 🧪 Sample Output

```
Integer: 42
Float: 42
String: 42
```

## 🧩 Summary

| Feature             | Description                                |
|---------------------|--------------------------------------------|
| Explicit Conversion | Requires manual type conversion            |
| Syntax              | `TypeName(value)`                          |
| Strict Typing       | Go doesn't allow automatic type changes    |