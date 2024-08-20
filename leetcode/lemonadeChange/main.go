package main

import "fmt"

func lemonadeChange(bills []int) bool {
	remaining := map[int]int{}
	for _, bill := range bills {
		switch bill {
		case 5:
			remaining[bill]++
		case 10:
			if remaining[5] > 0 {
				remaining[5]--
				remaining[10]++
			} else {
				return false
			}
		case 20:
			if remaining[10] > 0 && remaining[5] > 0 {
				remaining[10]--
				remaining[5]--
			} else if remaining[5] >= 3 {
				remaining[5] -= 3
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 10, 10, 20}))
	fmt.Println(lemonadeChange([]int{10, 10}))
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 20}))
}
