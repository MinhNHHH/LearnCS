package main

import "fmt"

func removeOccurrences(s string, part string) string {
	stack := []rune{}

	for _, char := range s {
		stack = append(stack, char)
		start := len(stack) - len(part)
		end := len(stack)
		if len(stack) >= len(part) && string(stack[start:end]) == part {
			stack = stack[:start]
		}
	}
	return string(stack)
}

func main() {
	fmt.Println(removeOccurrences("daabcbaabcbc", "abc"))
}
