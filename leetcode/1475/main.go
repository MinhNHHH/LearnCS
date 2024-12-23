package main

import "fmt"

func finalPrices(prices []int) []int {
	for i := 0; i < len(prices)-1; i++ {
		j := i + 1
		for j < len(prices) {
			if prices[i] >= prices[j] {
				prices[i] = prices[i] - prices[j]
				break
			}
			j++
		}
	}

	return prices
}
func main() {
	fmt.Println(finalPrices([]int{8, 4, 6, 2, 3}))
}
