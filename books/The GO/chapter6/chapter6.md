# 6.1 Method declaration

- Methods are special functions with a receiver argument.
- The receiver appears in its own arg list between the `func` kw and the method name.

```go
func (receiverType ReceiverName) MethodName(parameters) (returnType) {
    // method body
}
```

- `Receiver`: The receiver is a parameter that denotes which type the method belongs to.
- `Method name`: the name of method.
- `Parameters`: The list of parameters the method takes.
- `Return type`: The type of value the method returns.

## EXAMPLE

```go
package main

import (
    "fmt"
)

// Define a struct type
type Circle struct {
    Radius float64
}

// Define a method on the Circle type
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// Define a method with a pointer receiver
func (c *Circle) SetRadius(r float64) {
    c.Radius = r
}

func main() {
    c := Circle{Radius: 5}
    fmt.Println("Area:", c.Area())  // Call method with value receiver

    c.SetRadius(10)  // Call method with pointer receiver
    fmt.Println("New Radius:", c.Radius)
    fmt.Println("New Area:", c.Area())
}
```

### Value Receivers

- A value receiver is a method that operates on a copy of the value of the receiver. It does not modify the original value.

```go
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

- `Behavior`: Since the method operates on a copy, any modifications made to the receiver within method do not affect the original value.
- `Usage`:
  - `Immutable types`: when the method does not need to modify the receiver.
  - `Small structs`: when the receiver is a small struct, and the overhead of copying the value is minimal.
  - `Read-only-operations`: When the method is intended to be read-only and will not alter the receiver's state.

### Pointer Receivers

- A pointer receiver is a method that operates on a pointer to the receiver. It can modify the original value.

```go
func (r *Rectangle) Scale(factor float64) {
  r.Width *= factor
  r.Height *= factor
}
```

- `Behavior`: Since the method operates on a pointer, any modifications make to the receiver within the method affect to the original value.
- `Usage`:
  - `Mutable types`: When the method needs to modify the receiver.
  - `Large structs`: When the receiver is a large structs and copying the value would be expensive in terms of memory and performance.
  - `Efficient Memory Use`: When passing a pointer is more efficient than copying a large struct.

```go
package main

import "fmt"

type Rectangle struct {
    Width, Height float64
}

// Area method with value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
  rect := Rectangle{Width: 10, Height: 5}
  area := rect.Area()
  fmt.Println("Area:", area)
  // Original rectangle dimensions remain unchanged
  fmt.Println("Original Rectangle:", rect)
}
```

# 6.4 Method Values and Expressions.
- Methods are a key feature that allows you to define functions associated with types.
- Go supports method values and method expressions, which provide more flexible ways to use methods.

### Method Values
- A method values is a function that is bound to a specific receiver.
- This means you can capture a method from an instance of a type and store it as a variable, which you can then call like a regular function.

```go
package main

import "fmt"

type Rectangle struct {
    Width, Height float64
}

// Area method with value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    
    // Capture the method value
    areaFunc := rect.Area
    
    // Call the captured method
    fmt.Println("Area:", areaFunc())
}
```

### Method Expression
- A method expression is away to reference a method as a function without binding it to a specific receiver.


```go
package main

import "fmt"

type Rectangle struct {
    Width, Height float64
}

// Area method with value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    
    // Create a method expression
    areaFunc := Rectangle.Area
    
    // Call the method expression with rect as the receiver
    fmt.Println("Area:", areaFunc(rect))
}
```

### Differences between method values and method expressions.

- `Binding`: 
    - `Method Values`: Bind the method to a specific receiver instance. The receiver is fixed and does not need to be provided when calling the method.
    - `Method Expressions`: Do not bind the method to a specific receiver. The receiver must be provided explicitly when calling the method.

- `Use Cases`:
    - `Method Values`: useful when you want to capture the method along with its receiver and pass it around or store it for later use.
    - `Method Expressions`: Useful when you want a fucntion-like form of a method and need to call it with different receivers.
