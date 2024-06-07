// Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas.
// The guards have gone and will come back in h hours.

// Koko can decide her bananas-per-hour eating speed of k.
// Each hour, she chooses some pile of bananas and eats k bananas from that pile.
// If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.

// Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

// Return the minimum integer k such that she can eat all the bananas within h hours.

// Input: piles = [3,6,7,11], h = 8
// Output: 4

// Input: piles = [30,11,23,4,20], h = 5
// Output: 30

// Input: piles = [30,11,23,4,20], h = 6
// Output: 23

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
}

func minEatingSpeed(piles []int, h int) int {
	sort.Ints(piles)
	l, r := 1, piles[len(piles)-1]
	// Time complexity O(nlog(maxPiles))
	for l < r {
		mid := l + (r-l)/2

		count := 0
		for _, value := range piles {
			count += (value + mid - 1) / mid
		}

		if count <= h {
			r = mid
		} else {
			l = mid + 1
		}
	}
	// Space complexity (O(1))
	return l
}
