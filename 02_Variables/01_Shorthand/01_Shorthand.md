# Explanation

## Package and Import

- `package main`: Defines the executable program's entry point.

- `import "fmt"`: Imports the `fmt` package, which contains functions for formatted input/output.

## Variable Declarations

Variables are declared using the short declaration operator `:=`, which infers the variable type automatically.

- `a := 10` declares `a` as an integer with value 10.

- `b := "golang"` declares `b` as a string.

- `c := 4.17` declares `c` as a `float64` (floating point number).

- `d := true` declares `d` as a boolean.

- `e := "Hello"` declares `e` as a string.

- ``f := `Do you like my hat?` `` declares `f` as a raw string literal using backticks. This allows multi-line strings or strings with special characters without escaping.

- `g := 'M'` declares `g` as a rune, which is an alias for `int32`, representing a Unicode code point.

## Printing Variable Values

- `fmt.Printf("%v \n", var)` prints the value of the variable using the default format (`%v`).

- Each variable is printed on its own line.

## Printing Variable Types

- `fmt.Printf("%T \n", var)` prints the type of the variable.

- This helps you verify what Go inferred as the variable type.

## Sample Output

10
golang
4.17
true
Hello
Do you like my hat?
77

int
string
float64
bool
string
string
int32


- The rune `'M'` prints as `77`, its Unicode code point.

- The type for the rune is shown as `int32`.

## Summary

- Short variable declarations (`:=`) are a concise way to declare variables with automatic type inference.

- Variables of different types can be declared easily.

- You can print both the value and the type of any variable using `fmt.Printf`.

- Raw strings use backticks and are useful for unescaped, multi-line strings.

- Runes represent Unicode code points and use single quotes.
