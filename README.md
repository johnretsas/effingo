# effingo

A lightweight Go package that brings JavaScript-style array methods to Go slices.

## Installation

```bash
go get effingo
```

## Usage

```go
import "effingo/effingo"
```

### Map

Transform each element in a slice.

```go
doubled := effingo.Map([]int{1, 2, 3}, func(x int) int { return x * 2 })
// [2, 4, 6]
```

### Filter

Keep elements that match a condition.

```go
evens := effingo.Filter([]int{1, 2, 3, 4}, func(x int) bool { return x%2 == 0 })
// [2, 4]
```

### ForEach

Execute a function for each element.

```go
effingo.ForEach([]string{"a", "b", "c"}, func(s string) {
    fmt.Println(s)
})
```

### Reduce

Accumulate elements into a single value.

```go
sum := effingo.Reduce([]int{1, 2, 3, 4}, 0, func(x int, acc int) int { return acc + x })
// 10
```

### Find

Find the first element matching a condition. Returns the index and a pointer to the value (`nil` if not found).

```go
index, val := effingo.Find([]int{1, 2, 3, 4}, func(x int) bool { return x > 2 })
// index: 2, *val: 3
```
