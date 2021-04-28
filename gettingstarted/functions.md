**Functions**

Functions are built up in a "stack".

example 1.
```
func average(xs []float64) float64 {
    total := 0.0
        for _, v := range xs {
        total += v
    }
    return total / float64(len(xs))
}
```
**Returning Multiple Values**

Go is also capable of returning multi-values from a function.

```
func f() (int, int) {
    return 5, 6
}

func main() {
    x, y := f()
}
```

**Variadic Functions**

here is a special form available for the last parameter in a Go function:
```
func add(args ...int) int {
  total := 0
  for _, v := range args {
    total += v
  }
  return total
}
func main() {
  fmt.Println(add(1,2,3))
}

```
By using `...` before the type name of the last parameter 
you can indicate that it takes zero or more of those parameters.

We can also pass a slice of int by following the slice with ...:

```
func main() {
    xs := []int{1,2,3}
    fmt.Println(add(xs...))
}
```

**Closure**

It is possible to create function inside of functions.

```
func main() {
    add := func(x, y int) int {
        return x + y
      }
    fmt.Println(add(1,4))
}
```

**Defer, Panic & Recover**

`defer` schedules a function call to be run after the function completes.
It is often used when resource need to be freed in some way.
For example,
```
f, _ := os.Open(filename)
defer f.Close()
```
_3 Advantages_
1. it keeps `close` call near `open` call, easier to understand
2. if function had multi-return statement, `close` will happen before both of them
3. The deferred function are run even if a run-time panic occurs.

**Panic & Recover**

`panic` cause a run time error.
`recover` stops the panic and return the value that was passed to the call to `panic`

A panic generally indicates a programmer error 
for example,
- attempting to access an index of an array that's out of bounds, 
- forgetting to initialize a map, etc. or
- an exceptional condition that there's no easy way to recover from. 
