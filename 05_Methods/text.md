# Methods in Go

## ✅ What Are Methods?

Imagine you have a toy car. This toy car can **move**, **stop**, and **honk**. All these are **actions** that the car can do. In programming, we call these actions **methods**.

So, in Go, a **method** is a **function that belongs to a specific type** (like a struct).

> Think of a method as a **special function** attached to an object.

---

## 🤔 Why Do We Use Methods?

We use methods to:

* Group behavior (actions) with the data it works on.
* Make code cleaner and more organized.
* Work like real-world objects (car has a drive method, person has a speak method, etc.)

---

## 🧩 Syntax of a Method

```go
func (receiver ReceiverType) methodName(parameters) returnType {
    // method body
}
```

### Breakdown:

* `func`: keyword to define a function.
* `(receiver ReceiverType)`: special part that attaches the method to a **type** (like a struct).
* `methodName`: name of the method.
* `parameters`: optional input values.
* `returnType`: optional return value.

---

## 🧪 Example Code with Comments

```go
package main

import "fmt"

// This is our custom type (like a blueprint for a dog)
type Dog struct {
    Name string
    Age  int
}

// Method attached to Dog - it barks!
func (d Dog) Bark() {
    fmt.Println(d.Name, "says Woof!")
}

// Another method - dog celebrates birthday!
func (d *Dog) Birthday() {
    d.Age += 1
    fmt.Println(d.Name, "is now", d.Age, "years old!")
}

func main() {
    // Create a Dog value
    myDog := Dog{Name: "Buddy", Age: 3}

    // Call methods
    myDog.Bark()        // Output: Buddy says Woof!
    myDog.Birthday()    // Output: Buddy is now 4 years old!
}
```

### Notes:

* `func (d Dog) Bark()`:

  * `d` is a **copy** of the `Dog` value.
  * Used when you **don’t want to change** the original Dog.
* `func (d *Dog) Birthday()`:

  * `d` is a **pointer** to the `Dog` value.
  * Used when you **want to modify** the original Dog.

---

## 🌟 Key Points Summary

| Concept          | Description                                          |
| ---------------- | ---------------------------------------------------- |
| Method           | A function attached to a type (like struct).         |
| Receiver         | Gives the method access to the struct’s data.        |
| Value Receiver   | `(d Dog)` – gets a copy; doesn't change original.    |
| Pointer Receiver | `(d *Dog)` – gets reference; can modify original.    |
| Why use?         | To group actions with data, like real-world objects. |

---