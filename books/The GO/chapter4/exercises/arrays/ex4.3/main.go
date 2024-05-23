package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	reverse(&arr)
	fmt.Println(arr)
}

func reverse(s *[]int) []int {
	slice := *s // Dereference the pointer to get the slice
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
