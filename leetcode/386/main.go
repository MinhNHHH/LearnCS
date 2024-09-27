package main

import (
	"fmt"
)

func lexicalOrder(n int) []int {
	res := []int{}
	var helper func(num, n int)
	helper = func(num, n int) {
		if num > n {
			return
		}
		res = append(res, num)
		for i := 0; i <= 9; i++ {
			current := num*10 + i
			if current > n {
				return
			}
			helper(current, n)
		}
	}
	for i := 1; i <= 9; i++ {
		helper(i, n)
	}
	return res
}

func main() {
	fmt.Println(lexicalOrder(100))
}
