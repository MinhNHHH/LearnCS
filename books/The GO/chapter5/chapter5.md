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
