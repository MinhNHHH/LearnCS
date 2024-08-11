package main

import (
	"fmt"
)

func main() {
	var x int
	arr := [3]int{3, 5, 2}
	slice := arr[1:] // Slicing from index 1 to the end
	fmt.Println(x, arr)
}
