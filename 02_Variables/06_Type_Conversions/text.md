# 🔄 5. Type Conversion in Go

## 📘 Explanation

Go requires **explicit** type conversion. You must manually convert values from one type to another — there's no automatic (implicit) conversion like in some other languages.

## ✍️ Syntax

```go
converted := TypeName(original)
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