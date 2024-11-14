package main

import "fmt"

func maxArr(arr []int) int {
	maxV := 0
	for _, value := range arr {
		if maxV < value {
			maxV = value
		}
	}
	return maxV
}
func minimizedMaximum(n int, quantities []int) int {
	left := 1
	right := maxArr(quantities)

	for left < right {
		x := left + (right-left)/2
		storesNeed := 0
		for _, quan := range quantities {
			storesNeed += (x + quan - 1) / x
		}
		if storesNeed > n {
			left = x + 1
		} else {
			right = x
		}
	}
	return left
}

func main() {
	fmt.Println(minimizedMaximum(6, []int{11, 6}))
}
