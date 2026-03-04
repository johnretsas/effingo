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

There are two ways to use effingo: standalone functions that work on plain slices, or the `Array` wrapper for method chaining.

## Standalone Functions

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

## Array Wrapper

Wrap a slice with `From` to use method-style calls.

```go
arr := effingo.From([]int{1, 2, 3, 4, 5, 6})
```

### ForEach

```go
arr.ForEach(func(x int) {
    fmt.Println(x)
})
```

### Filter

```go
evens := arr.Filter(func(x int) bool { return x%2 == 0 })
```

### Find

```go
index, val := arr.Find(func(x int) bool { return x > 3 })
```

### MapTo

`MapTo` and `ReduceTo` are standalone functions since Go methods can't introduce new type parameters.

```go
doubled := effingo.MapTo(arr, func(x int) int { return x * 2 })
```

### ReduceTo

```go
sum := effingo.ReduceTo(arr, 0, func(x int, acc int) int { return acc + x })
```
