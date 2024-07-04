# Concurrency with Shared Variables.

## 9.1 Race Condition
- Race conditions occur in concurrent programming when multiple goroutines access shared resources without proper synchorination, leading to unpredictable and erroneous outcomes.

1. `Goroutines`: Lighweight threads managed by the Go runtime. They are used to perform concurrent tasks. When multiple goroutines access shared variables, race condition can occur if not properly synchronized.
2. `Shared Variables`: Variables accessed by multiple goroutines simultaneously. Without proper synchronization, the order of access and modifications not guanranteed, leading to inconsistent data.
3. `Atomic Operations`: Operations that complete in a single step relative other threads.
4. `Mutexes`: Mutual exclusion locks provided by the `sync` pacage to ensure that only one goroutine accesses a critical section of code at a time. Lock has two methods `Lock` and `Unlock`.
5. `Channels`: Go's primary concurrency construct that can be used to synchronize goroutines by sending and receiving values. Channels ensure safe communication between goroutines, avoiding the need for explicit locks.

- Example Run with Race Condition:
```go
package bank

import (
    "fmt"
    "sync"
)

var balance int

func Deposit(amount int) {
    balance += amount
}

func Balance() int {
    return balance
}

// Example usage of the unsafe bank package
func main() {
    var wg sync.WaitGroup
    wg.Add(3)

    // Concurrently depositing amounts
    go func() {
        Deposit(100)
        wg.Done()
    }()
    go func() {
        Deposit(200)
        wg.Done()
    }()
    go func() {
        Deposit(300)
        wg.Done()
    }()

    // Wait for all deposits to complete
    wg.Wait()

    // Checking the balance
    fmt.Printf("Current balance: %d\n", Balance())
}
```
- Let's see what might happen when the example is run without synchronization:

  + Goroutine 1: Reads balance (initially 0), adds 100, writes 100.
  + Goroutine 2: Reads balance (still 0), adds 200, writes 200.
  + Goroutine 3: Reads balance (could be 100 or 200), adds 300, writes the final value.
- Correct Output: Current balance: 600

- Correcting handler race condition:
```go
// Package bank provides a concurrency-safe bank with one account.
package main

import "fmt"

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

// Deposit adds amount to the balance.
func Deposit(amount int) { deposits <- amount }

// Balance returns the current balance.
func Balance() int { return <-balances }

func teller() {
    var balance int // balance is confined to teller goroutine
    for {
        select {
        case amount := <-deposits:
            balance += amount
        case balances <- balance:
        }
    }
}

// Init starts the teller goroutine.
func Init() {
    go teller() // start the monitor goroutine
}

// Example usage of the bank package.
func main() {
    Init()

    // Depositing amounts concurrently
    go Deposit(100)
    go Deposit(200)
    go Deposit(300)

    // Allow some time for deposits to be processed
    // (In a real application, use sync.WaitGroup or similar mechanism)
    fmt.Println("Deposited amounts")

    // Checking the balance
    fmt.Printf("Current balance: %d\n", Balance())
}
```

- Handler race condition with Mutex

```go
package main

import (
  "fmt"
  "sync"
)
var (
  balance int
  mu sync.Mutex
)

func deposit(amount int) {
  balance += amount
}

func Deposit(amount int) {
  mu.Lock()
  defer mu.Unlock()
  deposit(amount)
}

func Blance() int {
  mu.Lock()
  defer mu.Unlock()
  return balance
}

func Withdraw(amount int) bool {
  mu.Lock()
  defer mu.Unlock()

  deposite(-amount)

  if balance < 0 {
    deposite(amount)
    return false
  }
  return true
}
```

### RWMutex
-  Privdes a multiple readers, single writer lock. It allows multiple goroutines to read a resource concurrently, but only one goroutine can write to the resource at a time.

  + `Rlock`: Used by readers to acquire a read lock. Multiple goroutines can hold the read lock simultaneously.
  + `RUnlock`: Used by readers to release the read lock.
  + `Lock`: Used by a writer to acquire an exclusive write lock. Only one goroutine can hold the write lock, and no other goroutines can acquire a read or write lock while it's held.
  + `Unlock`: Used by a writer to releases the write lock
