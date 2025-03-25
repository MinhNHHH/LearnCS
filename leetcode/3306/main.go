package main

import "fmt"

func countOfSubstrings(word string, k int) int64 {
	res := 0
	if len(word) < k+5 {
		return int64(res)
	}

	for i := 0; i < len(word); i++ {

	}

	return int64(res)
}

func main() {
	fmt.Println(countOfSubstrings("aeioqq", 1))
}
