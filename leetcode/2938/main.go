package main

import "fmt"

func minimumSteps(s string) int64 {
	step := 0
	blackCount := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			step += blackCount
		} else {
			blackCount++
		}
	}
	return int64(step)
}

func main() {
	fmt.Println(minimumSteps("110"))
}
