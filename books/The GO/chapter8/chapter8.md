# 8.1 Goroutines

1. `Concurrency Model`: 
- Go promotes a concurrency based on goroutines and channels.
- Goroutines are functions or methods that run concurrently with other goroutines, not necessarily in parallel but managed efficiently.
2. `Lightweight`:
- Goroutines are lightweight in comparison to threads in traditional operating system.
- A typical Go program may have thousands or even millions of goroutines.

3. `Syntax`:
```go
func main() {
    go doSomething() // Start a new goroutine
    // Other main function logic
}
```

4. `Concurrency vs Parallelism`: 
- Goroutines enable concurrent programming, where multiple tasks can be in progress simultaneously.
- Go runtime manages how goroutines are scheduled on OS threads.

5. `Benefits`:
- Scalability: Goroutines make it easier to write efficient concurrent programs.