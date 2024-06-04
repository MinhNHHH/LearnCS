# 4.1 Slice

## 4.2.1 The append function [append-function](./arrays/append-funcion.go)

1.`Automatic Expansion`: if the slice's underlying array does not have enough capacity to accommodate the new elements, `append` create a new array, copies the existing elements, and then appends the new elements.

2.`Efficiency Considerations`: Although `append` is efficient, execessive use can lead to multiple allocations and copying. To minimize this, it's often useful to pre-allocate slices with an estimated capacity using `make`.

## 4.2.2 In-Place Slice Techniques

- Reusing the Input Slice

  - Instead of creating a new slice, we reuse the input slice for the result by slicing it to zero length. This allow us to append to the same underlying array, reducing memory allocation.

  ```go
   s := []int{1, 2, 3, 4, 5}
   s = s[:0] // Reslice to zero length, retaining the underlying array
  ```

- Removing Elements
  - Remove elements from a slice by reslicing and appending.

  ```go
   s := []int{1, 2, 3, 4, 5}
   i := 2
   s = append(s[:i], s[i+1:]...)
  ```

- Filtering Elements
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

- Squashing Runs of Values [squashing](../exercises/arrays/ex4.6/main.go)
  - Replace runs of consecutive values with a single value.

# 4.3 Maps

- The hash table is one of the most ingenious and versatile of all datastructures.
- It's an unordered collection of key/value pairs in which all the keys are distinct and the values associated with a given key can be retrieved, updated or removed using a constant number of key comparisons on the average.
- Go does not provide a `set` type, but since the keys of keys of a map are distinct, a map can server this purpose. `make(map[string]bool) // a set of strings`.

# 4.4 Struct

- A `struct` is an aggregate data type that groups together zero or more named values of arbitary types as a single entity.

## 4.4.1 Stuct Literals

- Definition: A struct is a composite data type that groups together variables under a single name. Structs are defined using the `type` keyword followed by the `strucut` keyword and a list of fields.
- Struct Literal Syntax: Struct literals provide a concise way to create instances of a struct type and initialize their fields.
- Example:

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

- Initialization: When using a struct literal, you specify the field names and their corresponding values within curly braces `{}`. The order of fields doesn't matter, as long as each field is assigned a value.
- Advantages:
  - Concise Syntax: Struct literals allow you to create and initialize structs in a single expression.
  - Flexibility: you can initialize only specific fields while leaving others with their zero values.
- Zero Values: If you don't specific a value for a field in a struct literal, it will be initialized with its zero value. (e.g.. 0 for numeric types, empty string for strings, nil for pointer)

## 4.4.2 Comparing Structs

- If all the fields of a struct are comparable, the struct itself is comparable, so two expressions of that type may be compared using `==` or `!=`.
- The `==` operation compares the corresponding fields of the two structs in order, so two printed expressions below are equivalent:

  ```go
  type Point struct{ X, Y int }
     p := Point{1, 2}
     q := Point{2, 1}
     fmt.Println(p.X == q.X && p.Y == q.Y) // "false"
     fmt.Println(p == q)                   // "false"
  ```

## 4.4.3 Struct Embedding and Anonymous Fields

### Struct Embedding

- Struct embedding allows one struct type to be embedded within another struct type.
- This is a form of composition and can be thought of as a way to inherit the fields and methods of the embedded struct.

```go
 type Person struct {
  Name string
  Age int
 }

 type Employee struct {
  Person
  ID int
  Position string
 }

 func main() {
  e := Employee {
   Person : Person {
    Name : "John",
    Age: 25,
   },
   ID: 1,
   Position: "Software Engineer",
  }
 }
```

### AnonymousFields

- Anonnymous fields are fields in a struct withou a name, which can only be other struct types or interfaces.
- Similar with Embeding Struct. Anonymous Fields allows direct access to the fields and methods of the anonymous fields.

```go
 type Address struct {
   City, State string
 }

 type Contact struct {
   Name    string
   Address // Anonymous field
 }

 func main() {
   c := Contact{
     Name: "Alice",
     Address: Address{
       City:  "Wonderland",
       State: "Dreamland",
     },
   }

   // Accessing fields of the anonymous struct directly
   fmt.Println(c.Name)  // Output: Alice
   fmt.Println(c.City)  // Output: Wonderland
   fmt.Println(c.State) // Output: Dreamland
 }
```

### Methods on Embedded Types

- If the embedded struct has methods, those methods are also promoted to the embedding struct.

```go
 type Animal struct {
  Name string
 }

 func (a Animal) Speak() {
  fmt.Println(a.name, "says hello")
 }

 type Dog struct {
  Animal
  Breed string
 }

 func main() {
  d := Dog{
   Animal: Animal {
    Name: "Buddy",
   },
   Breed: "Golden Retriever",
  }
 }

 d.Speak() // Output: Buddy says hello
```

```
 - Struct Embedding and anonymous fields in Go provide a way to create flexible and reusable code.
 - By embedding structs, you can inherit fields and methods, and with anonymous fields, you can simplify access to those fields and methods, promoting clean and concise code.
```

# 4.5 JSON

- JavaScript Object Notation (JSON) is a standard notation for sending and receiving structured information.
- A JSON array is an ordered sequenece of values, written as a comma-separated list enclosed in square bracket.
- A JSON object is a mapping from strings to values, written as a sequence of name:value pairs.

```go
 type Movie struct {
 Title string
 Year
 int `json:"released"`
 Color bool `json:"color,omitempty"`
 Actors []string
 }
 var movies = []Movie{
 {Title: "Casablanca", Year: 1942, Color: false,
 Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
 {Title: "Cool Hand Luke", Year: 1967, Color: true,
 Actors: []string{"Paul Newman"}},
 {Title: "Bullitt", Year: 1968, Color: true,
 Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
 }
```

1. Converting a Go data structure likes movies to JSON is called marshaling.

```go
 data, err := json.Marshal(movies)
 if err != nil {
  log.Fatalf("JSON marshaling failed: %s", err)
 }
 fmt.Printf("%s\n", data)
```

2. Decoding (Unmarshalling) JSON to GO structs.

```go
type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Email  string `json:"email"`
}

func main() {
	jsonData := []byte(`{"name":"John Doe","age":30,"email":"john.doe@example.com"}`)

	var person Person
	err := json.Unmarshal(jsonData, &person)
	if err != nil {
			log.Fatal(err)
	}

	fmt.Printf("%+v\n", person)
}
```

3. Handling Nested JSON

- Go structs can also be used to represent nested JSON structures by using nested strucks.
```go

type Address struct {
    Street string `json:"street"`
    City   string `json:"city"`
    State  string `json:"state"`
}
// Nested struct
type Person struct {
    Name    string  `json:"name"`
    Age     int     `json:"age"`
    Email   string  `json:"email"`
    Address Address `json:"address"`
}

func main() {
	jsonData := []byte(`{
		"name": "John Doe",
		"age": 30,
		"email": "john.doe@example.com",
		"address": {
			"street": "123 Main St",
			"city": "Anytown",
			"state": "CA"
		}
	}`)

	var person Person
	err := json.Unmarshal(jsonData, &person)
	if err != nil {
			log.Fatal(err)
	}

	fmt.Printf("%+v\n", person)
}
```

4. Working with Arbitrary JSON

For handling arbitrary JSON structures where you may not know the strucutre in advance, you can use `map[string]interface{}` or `interface{}`

```go
func main() {
	jsonData := []byte(`{
			"name": "John Doe",
			"age": 30,
			"email": "john.doe@example.com",
			"address": {
					"street": "123 Main St",
					"city": "Anytown",
					"state": "CA"
			}
	}`)

	var data map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
			log.Fatal(err)
	}

	fmt.Printf("%+v\n", data)
}
```

# 4.6. Text and HTML Templates.
