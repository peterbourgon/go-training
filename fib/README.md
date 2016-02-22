# fib

```go
func fibonacci(n int, c chan int) {
    // Implement this function
}

func main() {
    c := make(chan int)
    go fibonacci(10, c)
    for i := range c {
        println(i)
    }
}
```
