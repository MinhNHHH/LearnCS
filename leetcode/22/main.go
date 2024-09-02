package main

import (
	"fmt"
	"strings"
)

func generateParenthesis(n int) []string {
	res := []string{}
	stack := []string{}

	var helper func(openN, closedN int)
	helper = func(openN, closedN int) {

		if openN == closedN && openN == n {
			res = append(res, strings.Join(stack, ""))
			return
		}
		if openN < n {
			stack = append(stack, "(")
			helper(openN+1, closedN)
			stack = stack[:len(stack)-1]
		}
		if closedN < openN {
			stack = append(stack, ")")
			helper(openN, closedN+1)
			stack = stack[:len(stack)-1]
		}
	}

	helper(0, 0)

	return res
}

func main() {
	fmt.Println(generateParenthesis(3))
}
