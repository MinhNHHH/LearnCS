package main

import (
	"fmt"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	res := make([]int, len(spells))
	for index, spell := range spells {
		left, right := 0, len(potions)-1
		for left <= right {
			mid := (left + right) / 2
			if spell*potions[mid] >= int(success) {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		res[index] = len(potions) - left
	}
	return res
}

func main() {
	fmt.Println(successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7))
}
