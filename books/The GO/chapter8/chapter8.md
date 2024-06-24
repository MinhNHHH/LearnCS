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

# 8.4 Channels:
- A channel is a communication mechanism that lets one goroutine send values to another goroutine.
- Each channel is a conduit for values of a particular type, called the channel's element type.
- To create a channel, we use the built-in make function.

```go
ch := make(chan int)
ch := make(chan string)
ch := make(chan struct{})
```

- A channel has two principal operations, `send` and `receive` collectively known as communications.
    - A `send` statement transmits a value from one goroutine through the channel to another goroutine executing a correspoding `receive` expression.
    - Both operations are written using the `<-` operator.
    ```go
        ch <- x // a send sattement
        x = <- ch // a receive expression in an assignment statement
        <- ch // a receive statement, result is discarded.
    ```

- Channels support a third operation `close` which set a flag indicating that no more value will ever be sent on this channel.

```go
    close(ch)
```

## 8.4.1 Unbuffered Channels
- An unbuffered channel is an a channel that has nocapacity to store data.
- When a value is sent on an unbuffered channel,  the sending goroutine is blocked until another goroutine receives the value.
- This mean when a goroutine tries to receive a value from an unbuffered channel, it's blocked until another goroutine sends a value. This blocking behavior provides a natural synchronization mechanism between goroutines.

### Synchronization with Unbuffered Channels.
- Unbuffered channels are often used to synchronize goroutines.

```go
func main() {
    done := make(chan struct{})
    
    go func() {
        fmt.Println("Hello, World!")
        done <- struct{}{} // Signal completion
    }()
    
    <-done // Wait for signal
}
```

- Here, the `done` channel is used to synchronize the completion of the the anonymous goroutine with the main goroutine.
- The main go routine blocks on `<-done` until the anonymous goroutine sends a value ensuring that "Hello, World!" is printed before the program exits.

## 8.4.2 Pipelines.
