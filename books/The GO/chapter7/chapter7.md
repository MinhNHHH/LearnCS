# 7.1 Interfaces
- An interface in Go is a type that specifies a set of method signatures.
- Any type that implements all these method signatures is said to satisfy the interface.

### EXAMPLE
```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

## Implementing an interface

```go
type ByteCounter int

func (bc *ByteCounter) Write(p []byte) (int, error) {
  *bc += ByteCounter(len(p))
  return len(p), nil
}
```

### Practical Example

```go
package main

import (
  "fmt"
  "bytes"
  "bufio"
)

// Writer interface
type Writer interface {
  Write(p []byte) (n int, err error)
}

// WordCounter counts the number of words written to it.
type WordCounter int

// Write implements the Writer interface for WordCounter.
func (wc *WordCounter) Write(p []byte) (int, error) {
  scanner := bufio.NewScanner(bytes.NewReader(p))
  scanner.Split(bufio.ScanWords)
  count := 0
  for scanner.Scan() {
      count++
  }
  *wc += WordCounter(count)
  return len(p), scanner.Err()
}

// LineCounter counts the number of lines written to it.
type LineCounter int

// Write implements the Writer interface for LineCounter.
func (lc *LineCounter) Write(p []byte) (int, error) {
  scanner := bufio.NewScanner(bytes.NewReader(p))
  scanner.Split(bufio.ScanLines)
  count := 0
  for scanner.Scan() {
      count++
  }
  *lc += LineCounter(count)
  return len(p), scanner.Err()
}

func main() {
  text := "hello world\nthis is a test\nanother line"

  var wc WordCounter
  var lc LineCounter

  writeTo(&wc, text)
  writeTo(&lc, text)

  fmt.Printf("Word count: %d\n", wc) // Output: Word count: 6
  fmt.Printf("Line count: %d\n", lc) // Output: Line count: 3
}

func writeTo(w Writer, data string) {
  w.Write([]byte(data))
}

```

# 7.2. Interface Types
- An interface type specifies a set of method signatures.
- - When a type provides definitions for all the methods in an interface, it is said to implement that interface.
1. Interface Definition:
```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

2. Implicit Implement:
- A type implements an interface simply by providing implementations for the interface's methods.
- There is no explicit declaration of intent
```go
type MyWriter struct{}

func (mw MyWriter) Write(p []byte) (n int, err error) {
    // Implementation here
    return len(p), nil
}
```

3. Using Interfaces:
- Interfaces are often used to define functions that can accept any type that implements a particular interface.
```go
func writeData(w Writer, data []byte) error {
    _, err := w.Write(data)
    return err
}
```
4. Empty Interface
- The empty interface, `interface{}` is a special case. It can hold any value because it has no methods, making every type satisfy it.

```go
func printAnything(v interface{}) {
    fmt.Println(v)
}
```