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

# 7.4 Parsing Flags with flag.Value

- the `flag` package provides a convenient way to parse cmd args.
- When u need custom parsing for your flags, you can implement the `flag.Value`. This interface requires the following two methods:
  + `String() string`: This method returns the string representation of the flag's value.
  + `Set(value string) error` This method sets the flag's value from a string.

### Steps to implement custom flags
1. Define a Custom Type:
```go
  type Celsius float64
```

2. Create a struct for the custom flag:
- Define a struct that will implement the `flag.Value` interface, embedding the custom type.

```go
type celsiusFlag struct {
    Celsius
}
```
3. Implement the `flag.Value` interface:
Provide the `String` and `Set` methods for your struct.

```go
func (f *celsiusFlag) String() string {
    return fmt.Sprintf("%v°C", f.Celsius)
}

func (f *celsiusFlag) Set(s string) error {
    value, err := strconv.ParseFloat(s, 64)
    if err != nil {
        return err
    }
    f.Celsius = Celsius(value)
    return nil
}
```

4. Define a function to create the Flag:
- Create a function that registers the custom flag with the `flag` package.
```go
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
    f := celsiusFlag{value}
    flag.CommandLine.Var(&f, name, usage)
    return &f.Celsius
}
```

# 7.5 Interface values
- Interfaces area a fundamental part of the language that allows for defining behaivor without specifying concrete types.
- An interface type is defined by a set of method signatures.
- A type satisfies an interface if it implements all the methods declared in the interface.

### Basic Concepts of interfaces:
1. Defining an interface:
- An interface type is defined using the `type` keyword followed by a list of method signatures.
```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```
2. Implementing an Interface:
- Ay type impolements an interface by providing definitions for all the methods in the interface.
```go
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}
```

3. Using interface values:
- Interface values can hold any value that implements the interface.
- This allows for writing functions that operate on any type that satisfies the interface.

### Dynamic types and value.
  - A `dynamic type`: the concrete type of the value assigned to the interface.
  - A `dynamic value`: the actual value that satisfies the interface.

```bash 
  when an interface value is nil, both the dynamic type and value are nil
```

- For example, consider an interface `Shape` and type `Rectangle` that implements this interface:
```go
type Shape interface {
  Area() float64
  Perimeter() float64
}

type Rectangle struct {
  Width, Height float64
}

func (r Rectangle) Area() float64 {
  return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
  return 2 * (r.Width + r.Height)
}
var s Shape = Rectangle{Width: 10, Height: 5}
```
- When assign a `Rectangle`to a `Shape` variable, the interface value holds both the `dynamic type` and `the dynamic` value of the `Rectangle`
- Here `s` is an interface value where:
  - The `dynamic type` is `Rectangle`
  - The `dynamic value` is `Rectangle{Width: 10, Height: 5}`
