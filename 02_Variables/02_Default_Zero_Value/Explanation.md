# 🧠 Default Zero Values in Go

## 📦 Package and Import

- `package main`: Declares the executable entry point.
- `import "fmt"`: Imports the `fmt` package for formatted input/output operations.

---

## 🧮 Variable Declarations

The program declares variables without initializing them:

- `var a int`
- `var b string`
- `var c float64`
- `var d bool`

In Go, variables declared without an explicit initial value receive a **zero value** appropriate to their type:

| Variable | Type     | Zero Value     |
|----------|----------|----------------|
| `a`      | `int`    | `0`            |
| `b`      | `string` | `""` (empty)   |
| `c`      | `float64`| `0.0`          |
| `d`      | `bool`   | `false`        |

---

## 🖨️ Output Formatting

- `%v` in `fmt.Printf` is used to print each variable's value in its default format.
- An empty line appears for `b` since it's an empty string.

---


---

## 📌 Summary

- Go assigns a **default zero value** to variables that are declared but not initialized.
- This ensures all variables have predictable starting values.
- It helps avoid bugs caused by uninitialized data.


