package main

import "fmt"

func isValid(s string) bool {
	hash := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	stack := []string{}

	for _, char := range s {
		if len(stack) > 0 && stack[len(stack)-1] == hash[string(char)] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, string(char))
		}
	}

	if len(stack) > 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("([)"))
}
