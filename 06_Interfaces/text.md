## 🧩 What is an Interface in Go?

An **interface** in Go is a **type** that specifies a set of method signatures (behavior), but **not** their implementation.
If a type (like a struct) **implements all the methods** declared in an interface, it is said to **satisfy** that interface **implicitly**—without needing explicit declaration.

---

## 🎯 Why Use Interfaces?

* **Abstraction**: Interfaces allow you to write **generic code** that works with different types.
* **Decoupling**: You can **separate** the code that uses a type from the code that implements it.
* **Polymorphism**: Enables **different types to be treated the same** if they satisfy the same interface.
* **Testing**: You can **mock behaviors** using interfaces (great for unit testing).

---

## 🧪 Syntax

```go
type InterfaceName interface {
    MethodName(param1 Type1, param2 Type2) ReturnType
}
```

### Implementing the Interface:

Any type that has the same method **automatically implements** the interface.

```go
type MyStruct struct{}

func (m MyStruct) MethodName(param1 Type1, param2 Type2) ReturnType {
    // implementation
}
```

---

## 🧰 Example Code with Comments

```go
package main

import (
	"fmt"
)

// 1. Define the interface
type Speaker interface {
	Speak() string
}

// 2. Create a struct that implements the interface
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

// 3. Another struct that also implements the interface
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

// 4. Function that takes an interface as a parameter
func makeItSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	// Use the interface to hold different concrete types
	var s Speaker

	s = Dog{} // Dog satisfies Speaker
	makeItSpeak(s)

	s = Cat{} // Cat also satisfies Speaker
	makeItSpeak(s)

	// Or directly pass values
	makeItSpeak(Dog{})
	makeItSpeak(Cat{})
}
```

### 🧼 Output:

```
Woof!
Meow!
Woof!
Meow!
```

---

## 🗝️ Key Points

* Interfaces contain **only method signatures**, no data or implementation.
* **Satisfaction is implicit**—no need to declare `implements`.
* Interface variables can **hold any type** that implements the interface.
* You can create functions that **accept interfaces as parameters**, making them more reusable.
* Use `interface{}` to represent **any type** (also known as the empty interface).
* The **zero value** of an interface is `nil`.
* You can use **type assertions** or **type switches** to retrieve the concrete value.
