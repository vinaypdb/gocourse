# 🔁 for with range in Go

## 📘 Explanation

**🧠 Purpose**

for with range is used to **iterate over elements**  in:

- Arrays

- Slices

- Maps

- Strings

- Channels

It provides a clean and concise way to loop through elements **without needing to manually manage indices or conditions**.

**🤔 Why Use range?**

Use it when:

- You want to **access elements** in a collection easily.

- You want to **iterate over keys/values** in maps or characters in strings.

- You prefer **readable and safe looping**.

## 🧱 Syntax

```go
for index, value := range collection {
    // loop body
}

```
- `index` – position (or key, for maps)

- `value` – item at that position

You can **ignore values** using `_`.



## ✅ Example Code 1. Array or Slice

```go
nums := []int{10, 20, 30}

for i, val := range nums {
    fmt.Printf("Index: %d, Value: %d\n", i, val)
}

```

## 🧪 Sample Output

```go
Index: 0, Value: 10
Index: 1, Value: 20
Index: 2, Value: 30

```
## ✅ Example Code 2. Ignore index or value

```go
// Only value
for _, val := range nums {
    fmt.Println(val)
}

// Only index
for i := range nums {
    fmt.Println(i)
}

```

## 🧪 Sample Output

```go
Index: 0, Value: 10
Index: 1, Value: 20
Index: 2, Value: 30

```
## ✅ Example Code 3. Map

```go
m := map[string]int{"a": 1, "b": 2}

for key, val := range m {
    fmt.Printf("Key: %s, Value: %d\n", key, val)
}

```

## ✅ Example Code 4. String (Unicode aware)

```go
s := "Go語"

for i, r := range s {
    fmt.Printf("Index: %d, Rune: %c\n", i, r)
}


```


## 🔄 Common Mistakes Beginners Make

**1. Modifying the slice/map during range**
```go
for i, v := range slice {
    slice = append(slice, v) // ❌ can lead to unexpected behavior
}

```

**2. Reusing loop variables in goroutines**

```go
for i, v := range nums {
    go func() {
        fmt.Println(i, v) // ❌ might capture wrong i, v
    }()
}

```
✅ **Fix:**

```go
for i, v := range nums {
    go func(i, v int) {
        fmt.Println(i, v)
    }(i, v)
}

```
**3. Using range on nil slices or maps**

```go
var s []int
for _, v := range s {
    fmt.Println(v) // ✅ No panic — just skips
}

```
✅ Safe, but make sure it’s intended.

## 💡 Key Points

| Point                                   | Description                        |
| --------------------------------------- | ---------------------------------- |
| `range` simplifies iteration            | Over many data types               |
| Automatically gives index/key and value | No need for manual indexing        |
| Use `_` to ignore unused values         | e.g., `for _, v := range ...`      |
| Safe on empty or nil collections        | Just results in 0 iterations       |
| Be cautious inside goroutines           | Use argument capture to avoid bugs |



## 🧠 Summary

- `for ... range` is idiomatic in Go for looping through collections.

- It is safe, readable, and avoids off-by-one errors.

- Use it over traditional loops when working with arrays, maps, or strings.

- Watch out for closures and modifications during iteration.

