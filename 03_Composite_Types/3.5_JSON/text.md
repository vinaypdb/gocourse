# 📘 3.5 JSON

🧠 **What is JSON?**

**JSON JavaScript Object Notation)** is a lightweight data format used for exchanging data between a server and client, or between different parts of a program. It is:

- Human-readable

- Easy to generate and parse

- Supported by most languages (including Go)

**Why Use JSON in Go?**

- To **send and receive structured data** (e.g. APIs).

- To **serialize** (convert Go structs to JSON) and **deserialize** (convert JSON to Go structs).

- Often used in **web development, APIs, configuration files**, etc.

## ✍️ Syntax Overview

Go provides the `encoding/json` package for working with JSON.

- **Marshal:** Go struct → JSON

- **Unmarshal:** JSON → Go struct


## ✅ Example Code

```go
package main

import (
    "encoding/json"
    "fmt"
)

// Define a struct to match JSON structure
type Person struct {
    Name  string `json:"name"`  // JSON key: "name"
    Age   int    `json:"age"`   // JSON key: "age"
    Email string `json:"email"` // JSON key: "email"
}

func main() {
    // Create a struct instance
    person := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}

    // Marshal: Convert struct to JSON
    jsonData, err := json.Marshal(person)
    if err != nil {
        fmt.Println("Error marshalling:", err)
        return
    }

    // Print JSON string
    fmt.Println("JSON Output:", string(jsonData))

    // Unmarshal: Convert JSON back to struct
    var newPerson Person
    jsonStr := `{"name":"Bob","age":25,"email":"bob@example.com"}`

    err = json.Unmarshal([]byte(jsonStr), &newPerson)
    if err != nil {
        fmt.Println("Error unmarshalling:", err)
        return
    }

    // Print the struct
    fmt.Println("Struct Output:", newPerson)
}

```

## 🔎 Output

```
JSON Output: {"name":"Alice","age":30,"email":"alice@example.com"}
Struct Output: {Bob 25 bob@example.com}
```

## Key Points to Remember

- Use struct tags (`json:"key"`) to map struct fields to JSON keys.

- Use `json.Marshal()` to encode Go structs into JSON.

- Use `json.Unmarshal()` to decode JSON into Go structs.

- JSON keys must be **exported fields** (start with a capital letter) in the struct.

- Errors should always be checked when marshalling/unmarshalling.





