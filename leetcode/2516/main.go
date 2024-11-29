package main

import "fmt"

func takeCharacters(s string, k int) int {
	temp := map[rune]int{}

	for _, char := range s {
		temp[char]++
	}

	if temp[97] < k || temp[98] < k || temp[99] < k {
		return -1
	}

	left := map[rune]int{}
	right := map[rune]int{}
	left_sum, right_sum := 0, 0

	for _, char := range s {
		if left[char] == k {
			break
		}
		left[char]++
		left_sum++
	}

	for i := len(s) - 1; i >= 0; i-- {
		if right[rune(s[i])] == k {
			break
		}
		right[rune(s[i])]++
		right_sum++
	}

	if left_sum != right_sum {
		return left_sum + right_sum
	}

	return left_sum
}

func main() {
	fmt.Println(takeCharacters("a", 1))
	fmt.Println(takeCharacters("aabaaaacaabc", 2))
}
