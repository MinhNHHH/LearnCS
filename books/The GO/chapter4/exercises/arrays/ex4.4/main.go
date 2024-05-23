package main

import "fmt"

func rotate_ints(ints []int) {
	last := ints[len(ints)-1]
	first := ints[0]
	copy(ints, ints[:len(ints)-1])
	ints[0] = last
	ints[len(ints)-1] = first
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	rotate_ints(s)
	fmt.Println(s)
}
