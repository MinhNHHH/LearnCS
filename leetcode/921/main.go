package main

import "fmt"

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func minAddToMakeValid(s string) int {
	stack := []rune{}

	for _, char := range s {
		if len(stack) > 0 && string(stack[len(stack)-1]) == "(" && string(char) == ")" {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}

	}
	return len(stack)
}

func main() {
	fmt.Println(minAddToMakeValid("()))(("))
}
