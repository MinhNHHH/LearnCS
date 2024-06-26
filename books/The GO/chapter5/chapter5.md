# FUNCTION

+ A function lets us wrap up a sequence of statements as a unit that can be called from elsewhere in a program, perhaps multiple times.
+ Functions make it possible to break a big job into smaller pieces that might well be written by different people separeated by both time and space.
+ A function hides its implementetation details from its users.

## 5.1. Function Declarations
```go
  func nam(parameter-list) (result-list) {
    body
  }
```

- `parameter-list` specifies the names and types of the function's parameters, which are the local variables.
- `result-list` specifies the types of the values that the function returns. If the function returns one unnamed result or no results at all, parentheses are optional and ususally omitted.

## 5.2 Recursion
- Recursion is a programming technique where a function calls itself in order to solve a problem. Each recursive call works on a smaller or simpler sub-problem until a base condition is met, which stops the recursion.
- Key Concepts of Recursion:
  - Base Case: This is the condition that stops the recursion. Without a base case, the function would call itself indefinitely.
  - Recursive Case: This part of the fucntion reduces the problem into a smaller instance of the same problem and makes a recursive call.

```go
func factorial(n int) int {
	// Base case: if n is 0, return 1
	if n == 0 {
		return 1
	}
	// Recursive case: n! = n * (n-1)!
	return n * factorial(n-1)
}
```

## 5.3 Multiple Return Values
- A function can return more than one result.

```go
func traverseLink(n *html.Node, countWords, countImages int) (int, int) {
	if n == nil {
		return countWords, countImages
	}

	if n.Type == html.ElementNode && n.Data != "img" {
		countImages++
	} else if n.Type == html.TextNode {
		textList := strings.Split(n.Data, " ")
		for _, text := range textList {
			if isAlpha(text) {
				countWords++
			}
		}
	}

	countWords, countImages = traverseLinkz(n.FirstChild, countWords, countImages)
	countWords, countImages = traverseLinkz(n.NextSibling, countWords, countImages)
	return countWords, countImages
}

```

## 5.4 Erros
- Erros are fundamental mechanism for handling unexcepted situations or problems that arise during program execution.
- They provide a way to signal and communicate issues encountered by your code.
- `Error Interface`: The `error` interface is the foundation for representing erros.
```go
type error interface {
	Error() string
}
```
- `Built-in Errors`: Go provides serval built-in error types like `io.EOF` (end-of-file), `fmt.Errorf`(formatted error string creation),...
- Returning Erros: Function can return erros to indicate failures during their execution.
- `Error Handling`: There are two primary ways to handle errors in Go:
  - `if` Statements: We can use `if` statements to check for non-nil error values returned by functions. If an error is encountered, the program can take appropriate actions likes logging the error or exiting with error code.
  - `defer` statements: The `defer` statement allows you to execute code after the surrounding function returns, even if it panics or exits permaturely. This is commonly used to close resources (likes files) opened during the fucntion's execution, regardless of errors.
	```go
		// Example for `if` statement
		data, err := readFile("data.txt")
		if err != nil {
				fmt.Println("Error reading file:", err)
				return
		}
	```
	```go
	// Example for `defer` statement
		file, err := os.Open("data.txt")
		defer file.Close()  // Close file even on errors
		if err != nil {
				// Handle error
		}
	```
## 5.5 Function Values
- In go, functions are treated as first-class citizens, meaning they can be assigned to variables, passed around as arguments, and even returned from other functions.

1. Functions as Value:
  - Go allows you to assign a function to a variable. This variable then holds the functions value,.
2. Passing Functions As Arguments:
  - You can pass functions as arguments to other functions. This enables you to create functions that operate on other functions, providing more generic and higher-order programming capabilities.
3. Returning Functions from Functions:
  - In go, functions can also return functions. This allows you to create functions that generate or return other functions based on certain critera.

```go
package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func compute(fn func(int, int) int, a, b int) int {
	return fn(a, b)
}

func main() {
	// Assign function add to variable 'operation'
	operation := add

	// Pass 'operation' function as argument to 'compute'
	result := compute(operation, 5, 3)
	fmt.Println(result) // Output: 8
}
```


## 5.6 Anonymous function
- Anonymous functions are functions that are defined without a name.
- These functions are useful when you want to define a function on the fly and use it immediately.

### Defining and using anonymous function:
```go
package main

import "fmt"

func main() {
    // Defining an anonymous function and calling it immediately
    func() {
        fmt.Println("Hello from anonymous function!")
    }()
}
```

### Assigning Anonmous Functions to Variables
```go
package main

import "fmt"

func main() {
    // Assigning an anonymous function to a variable
    greet := func(name string) {
        fmt.Printf("Hello, %s!\n", name)
    }

    // Calling the anonymous function
    greet("Alice")
    greet("Bob")
}
```

### Anonymous Functions as Closures
```go
package main

import "fmt"

func main() {
    // Defining a function that returns an anonymous function (closure)
    counter := func() func() int {
        count := 0
        return func() int {
            count++
            return count
        }
    }

    // Using the closure
    increment := counter()
    fmt.Println(increment()) // Output: 1
    fmt.Println(increment()) // Output: 2
    fmt.Println(increment()) // Output: 3
}
```

### Passing Anonymous functions as Arguments
```go
package main

import "fmt"

// A function that takes another function as an argument
func applyOperation(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}

func main() {
    // Using anonymous functions as arguments
    sum := applyOperation(3, 4, func(x, y int) int {
        return x + y
    })
    product := applyOperation(3, 4, func(x, y int) int {
        return x * y
    })

    fmt.Println("Sum:", sum)         // Output: Sum: 7
    fmt.Println("Product:", product) // Output: Product: 12
}
```

## 5.7 Variadic Functions

- A variadic function is one that can be called with varying numbers of arguments.
```go
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
			total += val
	}
	return total

	fmt.Println(sum())           //  "0"
	fmt.Println(sum(3))          //  "3"
	fmt.Println(sum(1, 2, 3, 4)) //  "10"
}
```

- Although the ...int parameter behaves like a slice within the function body, the type of a variadic.
- Variadic functions are often used for string formatting.

## 5.8 Deferred Function calls
- The `defer` statement in Go is a powerful feature used to ensure that certain statements are executed when a function completes.
- It's commonly used for resource cleanup tasks, such as closing files, releasing locks, or any other operations that need to happen regardless of how the function exist.

#### How `derfer` work
- The `defer` statement is followed by a function call
- Execution Timing: The deferred function is not executed immediately but is scheduled to be run after the surrounding function returns.
- LIFO Order: If multiple `defer` statements are present, they are executed in LAST-IN-FIRST-OUT.

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("Start")
    
    defer fmt.Println("Deferred: 1")
    defer fmt.Println("Deferred: 2")
    defer fmt.Println("Deferred: 3")

    fmt.Println("End")
}
```


## 5.9 Panic
