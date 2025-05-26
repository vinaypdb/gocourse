# 🔄 Type Declarations in Go?

## 📘 Explanation

Imagine you're organizing a toy box.

- Some toys are **balls** 🎾

- Some are **cars** 🚗

- Some are **action figures** 🧍

Instead of calling every item just a “toy,” you give each type of toy a name: **Ball, Car, Figure.**

That’s what **type declaration** in Go is like.

You're telling Go:

"Hey, I want to **create my own kind of data type** and give it a name."

## 🧱 Why Use Type Declarations?

- It makes code easier to read.

- You can describe real-world things better.

- You can group related data together.

## 🧩 Two Simple Types of Type Declarations


**1.Create a new name for an existing type**
```go
type Age int
```
👶 This says: "I want to call the int type **Age** so I can use it for things like someone's age."
Now instead of writing:
```go
var myAge int = 20
```
You can write:
```go
var myAge Age = 20
```
It still behaves like an integer but looks more meaningful!



**2.Create a struct (a custom data type with multiple parts)**

Imagine you want to create a "person" with a **name** and **age**.

```go
type Person struct {
    name string
    age  int
}

```
🎒 This is like packing both `name` and `age` into one box called `Person`.


## ✅ Example Code

```go
package main

import "fmt"

// Step 1: Make your own type based on int
type Age int

// Step 2: Make a struct with multiple values
type Person struct {
    name string
    age  Age
}

func main() {
    var myAge Age = 22

    // Create a Person with name and age
    p := Person{name: "Alice", age: myAge}

    fmt.Println("My Age:", myAge)
    fmt.Println("Person Info:", p)
}


```


## 🧪 Sample Output

```
My Age: 22
Person Info: {Alice 22}

```

## 🧩 Key Points:

| 🧩 Term   | 🎯 Meaning                               |
| --------- | ---------------------------------------- |
| `type`    | Used to create a new type name           |
| `Age int` | `Age` is now just another name for `int` |
| `struct`  | Groups multiple values into one type     |
| `Person`  | A type that holds both `name` and `age`  |
