package main

import "fmt"

func minLength(s string) int {
	if len(s) == 1 {
		return 1
	}

	stack := []rune{}

	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && ((s[i] == 66 && stack[len(stack)-1] == 65) || (s[i] == 68 && stack[len(stack)-1] == 67)) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, rune(s[i]))
		}
	}
	fmt.Println(string(stack), stack)
	return len(stack)
}

func main() {
	fmt.Println(minLength("ACBBD"))
}
