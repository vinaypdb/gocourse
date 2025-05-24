# 📦 Custom Packages in Go

In Go, **packages** are the way to organize and reuse code. Every Go file belongs to a package. This guide walks you through the concept of creating and using **custom packages**.

---

## 🧠 What is a Package?

A **package** in Go is a collection of Go source files in the same directory that are compiled together. Go has two types of packages:

- **Executable packages**: These have a `main` package and produce an executable when compiled.
- **Library packages**: These define reusable code and are imported by other packages.

---

## 🛠 Why Use Custom Packages?

Using custom packages helps:
- **Organize** code into logical modules
- **Encapsulate** related functionality
- **Reuse** code across multiple files or projects
- **Maintain** large codebases easily

---

## 🧱 How to Create a Custom Package

### Step 1: Define a package in its own folder

**File:** `01_Getting_Started/greet/greet.go`

```go
package greet

import "fmt"

// Hello prints a greeting
func Hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
