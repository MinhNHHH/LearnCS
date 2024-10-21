package main

import (
	"fmt"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxUniqueSplit(s string) int {
	visited := map[string]bool{}

	var backTrack func(start int) int
	backTrack = func(start int) int {
		maxSplit := 0
		if start == len(s) {
			return 0
		}
		sub := ""
		for i := start; i < len(s); i++ {
			sub += string(s[i])
			if !visited[sub] {
				visited[sub] = true
				maxSplit = max(maxSplit, 1+backTrack(i+1))
				delete(visited, sub)
			}
		}
		return maxSplit
	}

	return backTrack(0)
}

func main() {
	fmt.Println(maxUniqueSplit("addbsd"))
}
