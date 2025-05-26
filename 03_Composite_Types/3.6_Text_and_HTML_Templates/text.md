# 📘 3.6 JSON

🧠 **What Are Templates in Go?**

Templates in Go allow you to **generate formatted output** using **dynamic data**. This is especially useful for:

Generating text reports (with `text/template`)

Rendering dynamic HTML web pages (with `html/template`)

**Why Use Templates?**

- To **separate logic from presentation**

- To **render structured output** (text or HTML) based on Go data

- Essential for building **web applications, emails, reports,** etc.

## ✍️ Types of Templates in Go

| Package         | Purpose               |
| --------------- | --------------------- |
| `text/template` | For plain text output |
| `html/template` | For safe HTML output  |



## ✅ 1.Text Template Example

```go
package main

import (
    "os"
    "text/template"
)

func main() {
    // Template with placeholder {{.}}
    tmpl := template.Must(template.New("greet").Parse("Hello, {{.}}!\n"))

    // Render the template with string data
    tmpl.Execute(os.Stdout, "Go Developer")
}

```

## 🔎 Output

```
Hello, Go Developer!

```
## ✅ 2.HTML Template Example

```
package main

import (
    "html/template"
    "net/http"
)

type User struct {
    Name  string
    Email string
}

func handler(w http.ResponseWriter, r *http.Request) {
    user := User{Name: "Alice", Email: "alice@example.com"}

    tmpl := template.Must(template.New("profile").Parse(`
        <html>
        <body>
            <h1>Hello, {{.Name}}</h1>
            <p>Email: {{.Email}}</p>
        </body>
        </html>`))

    tmpl.Execute(w, user)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

```
## 🔎 Output

```
<html>
<body>
    <h1>Hello, Alice</h1>
    <p>Email: alice@example.com</p>
</body>
</html>

```
Open http://localhost:8080 in a browser to see the result.

## Key Points to Remember

- `template.New(name).Parse(string)` creates a new template.

- `template.Must` ensures template creation errors are caught early.

- `tmpl.Execute(writer, data)` renders the template with data.

- `text/template` is for plain text; `html/template` escapes HTML to prevent XSS.

- Templates support `{{.Field}}` for accessing data fields and control structures like `if`, `range`, etc.





