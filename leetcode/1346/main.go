package main

import (
	"fmt"
)

func checkIfExist(arr []int) bool {
	visited := map[int]bool{}
	for _, num := range arr {
		if visited[num*2] {
			return true
		}
		if num%2 == 0 && visited[num/2] {
			return true
		}
		visited[num] = true
	}
	return false
}

func main() {
	fmt.Println(checkIfExist([]int{10, 2, 5, 3}))
}
