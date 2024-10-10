package main

import (
	"fmt"
	"sort"
)

func canArrange(arr []int, k int) bool {
	sort.Ints(arr)
	fmt.Println(((7 % 5) + 5) % 5)
	return true
}

func main() {
	fmt.Println(canArrange([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
}
