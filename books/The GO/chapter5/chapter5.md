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
