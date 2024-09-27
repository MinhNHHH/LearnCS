package main

import (
	"fmt"
	"strconv"
)

func diffWaysToCompute(expression string) []int {

	var helper func(left, right int) []int
	operator := map[string]func(a, b int) int{
		"-": func(a, b int) int { return a - b },
		"+": func(a, b int) int { return a + b },
		"*": func(a, b int) int { return a * b },
	}
	helper = func(left, right int) []int {
		res := []int{}
		for i := left; i <= right; i++ {
			op := string(expression[i])
			if f, ok := operator[op]; ok {
				num1 := helper(left, i-1)
				num2 := helper(i+1, right)
				for _, n1 := range num1 {
					for _, n2 := range num2 {
						result := f(n1, n2)
						res = append(res, result)
					}
				}
			}
		}
		if len(res) == 0 {
			num, _ := strconv.Atoi(expression[left : right+1])
			res = append(res, num)
		}
		return res
	}

	return helper(0, len(expression)-1)
}
func main() {
	fmt.Println(diffWaysToCompute("1-2-3"))
}
