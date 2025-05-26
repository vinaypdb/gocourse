# 🔄 Zero Values in Go

## 📘 Explanation

In Go, when you declare a variable **without initializing** it, Go automatically assigns it a **default "zero value"** depending on its type.

This ensures that all variables are initialized and there's **no concept of "uninitialized" variables** like in some other languages.

## ✍️ Syntax

```go
var a int        // zero value is 0
var b float64    // zero value is 0.0
var c string     // zero value is ""
var d bool       // zero value is false
```

## ✅ Example Code

```go
package main

import "fmt"

func main() {
    var a int
    var b float64
    var c string
    var d bool

    fmt.Println("int:", a)
    fmt.Println("float64:", b)
    fmt.Println("string:", c)
    fmt.Println("bool:", d)
}

```

## 🧪 Sample Output

```
int: 0
float64: 0
string: 
bool: false



```

## 🧩 Summary

| Feature             | Description                                |
|---------------------|--------------------------------------------|
| `int` |`0`|
| `float64`            | 	`0.0` |
| `string`    |`""` (empty string)  |
| bool   | false |


| Feature             | Description                                |
|---------------------|--------------------------------------------|
| Automatic Init |Variables are auto-initialized|
| Type-Specific	         | Each type has its own default zero value|
| Safe Defaults	     | Avoids undefined behavior or garbage values|
