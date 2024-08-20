package main

import "fmt"

func helper(n int, store map[int]int) int {
	if n <= 2 {
		return n
	}
	if value, ok := store[n]; ok {
		return value
	}
	store[n] = helper(n-1, store) + helper(n-2, store)
	return store[n]
}
func climbStairs(n int) int {
	store := map[int]int{}
	return helper(n, store)
}

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(45))
}
