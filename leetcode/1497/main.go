package main

import (
	"fmt"
)

func canArrange(arr []int, k int) bool {
	hash := map[int]int{}

	for _, num := range arr {
		key := ((num % k) + k) % k
		hash[key]++
	}
	for key := range hash {
		other := k - key
		if _, ok := hash[other]; ok {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(canArrange([]int{1, 2, 3, 4, 5, 6}, 10))
}
