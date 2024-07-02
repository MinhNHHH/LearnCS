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

### When to `close` a channel
1. Signaling completion:
- Close a channel to signal that no more data will be send.
- This is useful when the receiving goroutine needs to know that all data has been received.
```go
jobs := make(chan int, 5)
// Sender goroutine
go func() {
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs) // Signal no more jobs will be sent
}()

// Receiver goroutine
go func() {
    for job := range jobs {
        fmt.Println("Received job:", job)
    }
    fmt.Println("All jobs received")
}()
```
2. Broadcasting done signal:
- Close a channel to broadcast a done signal to multiple receivers.
- This is useful in case where multiple goroutines are waiting for a signal to stop working.
```go
done := make(chan struct{})

// Worker goroutine
go func() {
    <-done
    fmt.Println("Received done signal")
}()

// Closing the done channel signals all waiting goroutines
close(done)
```

3. Exiting Loops Gracefully:
- Close a channel when you need to exit a loop that depends on receiving values from that channel.
```go
data := make(chan int)

go func() {
    for i := 0; i < 10; i++ {
        data <- i
    }
    close(data)
}()

for num := range data {
    fmt.Println(num)
}
fmt.Println("No more data")
```

### When `not to close` a channel
1. Multiple seners:
- Do not close a channel if there are multiple concurrent senders.
- Only a signle sender should close the channel to avoid race conditions.
2. Shared Channels:
- Do not close a channel that you did not create unless you are sure that you are the only sender.
- This ensures that other goroutines are not unexpectedly disrupted.

### Guidelines for closing channels.
1. Responsibility:
    - The goroutine that creates that channel is usually responsible for closing it.
    - This prevents race conditions and ensures clear ownership
2. Idempotency: 
    - Closing a channel is idempotent.
    - Multiple calls to `close` will panic, so it should only be closed once.
3. Panic Handling:
    - Be aware that sending on a closed channel will cause a panic.
    - Ensure that no more sends are attempted after closing.

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
- Channels can be used to connect goroutines together so that the output of one is the input to another.
- Counter -> Squarer -> Printer
- The first goroutine `counter` generates the integers 0,1,2,3,... and sends them orver a channel to the second goroutine, `sqarer`, which receives each value, squares it, and sends the result over another channel to the third goroutine `printer` which receives the squared values and prints them.
- If the sender knows that no further values will ever be sent on a channel, it is useful to communicate this fact to the receiver goroutines so that they can stop waiting. This is accomplished by `closing` the channel use `close(ch)`.

## 8.4.3 Unidirectional
- We will break up large function in `gopl.io/ch8/pipeline2` above nito smaller pieces.

```go
    func counter(out chan int)
    func squarer(out, in chan int)
    func printer(in chan int)
```

- The `squarer` function sitting in the middle of the pipele, take two parameters, the input channel and the output channel. Bothe have the save type, but their intended uses are opposite: `in` is only to be received from, and `out` is only to be sent to.

-The Go type system provides unidirectional channel types that expose only one or the other of the send and receive operations.
```go
    chan<- int // send-only channel of int.
    <-chan int // receive-only channle of int.
```

## 8.4.4 Buffered Channels
- Buffered channels in go are a powerful feature that aloows goroutines to communicate with each other.
- Unlinke unbuffered channels, which block the sending goroutine until another goroutine receives from the channel, buffered channels provide a way to send and receive data without immediate synchronization between sender and receiver.

### What is buffered channels?
- A buffered channel has a capacity, defined when the channel si created.
- This capacity determines the number of elements that can be stored in the channel at any given time before any send operation blocks.

```go
    ch := make(chan int, e) // create buffered channel
    ch <- 1 // does not block
    ch <- 2 // does not block
    ch <- 3 // does not block
    ch <- 4 // blocks until space is available
```

### Advantages of buffered channels:
1. `Decoupling senders and receivers`: Buffered channels allow senders and receivers to run at different speeds by providing a buffer where values can be stored temporarily.
2. `Reducing synchronization overhead`: By allowing send operations to complete without blocking, buffered channels can reduce the overhead of synchronization in some scenarios.


### sync.WaitGroup
- It's a synchronization primitive used to manage the lifecycle and coordination of multiple goroutines.
1. Goroutine Synchronization:
- `sync.WaitGroup` allows a program to wait for a collection of goroutines to complete their execution.
2. Counter Mechanism:
    - Add: Increase of decreases the internal counter by delta. This is used to indicate the number of goroutines to wait for
    - Done(): Decreases the counter by one. This is typically called using `defer` in each go routines to signal its completion.
    - Wait(): Blocks until the counter becoimes zero, ensuring all goroutines have finished their work.

#### The goals of WaitGroup
1. Synchronize groutine comppletion:
    - Waiting for multiple goroutines
    - Tracking active goroutines.
2. Coordination of Concurrent Operations
    - Ensuring orderly execution
    - Avoiding premature exit

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure the counter is decremented when the goroutine completes
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c *net.TCPConn) {
	var wg sync.WaitGroup // Create a new WaitGroup for each connection
	input := bufio.NewScanner(c)
	for input.Scan() {
		wg.Add(1)
		go echo(c, input.Text(), 1*time.Second, &wg)
	}
	wg.Wait() // Wait for all echo goroutines to finish for this connection
	c.CloseWrite() // Close the write half of the connection
	c.Close() // Fully close the connection
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		tcpConn := conn.(*net.TCPConn)
		go handleConn(tcpConn)
	}
}
```

- In this example. So we implemented WaitGroup for each connection.
- Each connection should have its own `WaitGroup` because `WaitGroup` is used to manage and synchronize the lifecycle of goroutines specific to that connection.

Some reasons why having a separeate `WaitGroup`:
1. Isolation of Concurrency Management:
- `Separate Lifecycle Management`: Each connection can spawn multiple goroutines to handle its own each operations. Using a single `WaitGroup` for all conntions would mean that the `WaitGroup` is tracking goroutines across multiple connections, making it impossible to close individual connections correctly when their respective goroutines complete.
- `Correct Resource Cleanup`: each connection needs to ensure its resources are properly cleaned up when its goroutines sinsh. If all connections share the same `WaitGroup` the closure of one connection might be delayed until all aother connection's goroutines finsish, leading to resource leaks or incorrect behavior.

2. Accurate Synchronization.
- `Per-connection Synchronzation`: Each connection should wait for its own goroutines to finsh before closing the write half of the TCP connection. Sharing a `WaitGroup` across connections would make it challenging to determine when it's safe to cose specific connection since the `WaitGroup` would be waiting for goroutines from multiple connections to finsh.
- `Avoiding Interference`: If a `WaitGroup` is shared, any delay or long-running goroutine from one connection cloud block the closure of the other connections, causing interference and potentially leading to resource contention and degraded performance.

