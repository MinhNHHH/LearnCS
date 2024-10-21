package main

import "fmt"

func maximumSwap(num int) int {
	// Convert number into a slice of digits
	listNums := []int{}
	tempNum := num
	for tempNum > 0 {
		listNums = append(listNums, tempNum%10)
		tempNum = tempNum / 10
	}

	// Reverse listNums to represent the digits in the correct order
	for i, j := 0, len(listNums)-1; i < j; i, j = i+1, j-1 {
		listNums[i], listNums[j] = listNums[j], listNums[i]
	}

	// Create an array to store the maximum digit's index from the current index to the end
	n := len(listNums)
	maxIdx := make([]int, n)
	maxIdx[n-1] = n - 1
	// Traverse the digits array and store the max digit index from the right
	for i := n - 2; i >= 0; i-- {
		if listNums[i] > listNums[maxIdx[i+1]] {
			maxIdx[i] = i
		} else {
			maxIdx[i] = maxIdx[i+1]
		}
		fmt.Println(maxIdx)
	}

	// Try to find the first position to swap
	for i := 0; i < n; i++ {
		if listNums[i] != listNums[maxIdx[i]] {
			// Swap the current digit with the largest digit on the right
			listNums[i], listNums[maxIdx[i]] = listNums[maxIdx[i]], listNums[i]
			break
		}
	}

	// Rebuild the number from the modified digits
	res := 0
	for i := 0; i < n; i++ {
		res = res*10 + listNums[i]
	}

	return res
}

func main() {
	fmt.Println(maximumSwap(9976))
}
