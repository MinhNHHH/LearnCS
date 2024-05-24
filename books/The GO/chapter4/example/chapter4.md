# 4.1 Slice
## 4.2.1 The append function [append-function](./arrays/append-funcion.go)
  1.`Automatic Expansion`: if the slice's underlying array does not have enough capacity to accommodate the new elements, `append` create a new array, copies the existing elements, and then appends the new elements.

  2.`Efficiency Considerations`: Although `append` is efficient, execessive use can lead to multiple allocations and copying. To minimize this, it's often useful to pre-allocate slices with an estimated capacity using `make`.

## 4.2.2 In-Place Slice Techniques

+ Reusing the Input Slice
	- Instead of creating a new slice, we reuse the input slice for the result by slicing it to zero length. This allow us to append to the same underlying array, reducing memory allocation.
	```go
		s := []int{1, 2, 3, 4, 5}
		s = s[:0] // Reslice to zero length, retaining the underlying array
	```

+ Removing Elements
	- Remove elements from a slice by reslicing and appending.
	```go
		s := []int{1, 2, 3, 4, 5}
		i := 2
		s = append(s[:i], s[i+1:]...)
	```
+ Filtering Elements
	- Filter elements in place by iterating over the slice and keeping only the desired elements.
	```go
		s := []int{1, 2, 3, 4, 5}
		n := 0
		for _, x := range s {
			if x%2 == 0 {
				s[n] = x
				n++
			}
		}
		s = s[:n] // Reslice to include only the kept elements
	```
+ Squashing Runs of Values [squashing](../exercises/arrays/ex4.6/main.go)
	- Replace runs of consecutive values with a single value.
   
# 4.3 Maps
+ The hash table is one of the most ingenious and versatile of all datastructures.
+ It's an unordered collection of key/value pairs in which all the keys are distinct and the values associated with a given key can be retrieved, updated or removed using a constant number of key comparisons on the average.
+ Go does not provide a `set` type, but since the keys of keys of a map are distinct, a map can server this purpose. `make(map[string]bool) // a set of strings`.

# 4.4 Struct
+ A `struct` is an aggregate data type that groups together zero or more named values of arbitary types as a single entity.
## 4.4.1 Stuct Literals
+ Definition: A struct is a composite data type that groups together variables under a single name. Structs are defined using the `type` keyword followed by the `strucut` keyword and a list of fields.
+ Struct Literal Syntax: Struct literals provide a concise way to create instances of a struct type and initialize their fields.
+ Example:
	```go
	type Person struct {
    Name string
    Age  int
	}

	func main() {
		// Creating a struct instance using a struct literal
		person := Person{Name: "Alice", Age: 30}

		// Accessing struct fields
		fmt.Println("Name:", person.Name)
		fmt.Println("Age:", person.Age)
	}
	```
+ Initialization: When using a struct literal, you specify the field names and their corresponding values within curly braces `{}`. The order of fields doesn't matter, as long as each field is assigned a value.
+ Advantages:
	+ Concise Syntax: Struct literals allow you to create and initialize structs in a single expression.
	+ Flexibility: you can initialize only specific fields while leaving others with their zero values.
+ Zero Values: If you don't specific a value for a field in a struct literal, it will be initialized with its zero value. (e.g.. 0 for numeric types, empty string for strings, nil for pointer)
