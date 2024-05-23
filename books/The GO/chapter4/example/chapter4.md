## 4.2.1 The append function [append-function](append-funcion.go)
  1.`Automatic Expansion`: if the slice's underlying array does not have enough capacity to accommodate the new elements, `append` create a new array, copies the existing elements, and then appends the new elements.

  2.`Efficiency Considerations`: Although `append` is efficient, execessive use can lead to multiple allocations and copying. To minimize this, it's often useful to pre-allocate slices with an estimated capacity using `make`.

## 4.2.2 In-Place Slice Techniques

  1. Reusing the Input Slice
		- Instead of creating a new slice, we reuse the input slice for the result by slicing it to zero length. This allow us to append to the same underlying array, reducing memory allocation.
		```go
			s := []int{1, 2, 3, 4, 5}
			s = s[:0] // Reslice to zero length, retaining the underlying array
		```

  2. Removing Elements
       - Remove elements from a slice by reslicing and appending.
		```go
			s := []int{1, 2, 3, 4, 5}
			i := 2
			s = append(s[:i], s[i+1:]...)
		```
  3. Filtering Elements
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
  4. Squashing Runs of Values [squashing](../exercises/arrays/ex4.6/main.go)
       - Replace runs of consecutive values with a single value.
   
