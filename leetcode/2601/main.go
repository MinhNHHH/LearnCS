package main

import "fmt"

func sive(limit int) []int {
	isPrime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= limit; i++ {
		if isPrime[i] {
			for j := i * i; j <= limit; j++ {
				isPrime[j] = false
			}
		}
	}

	primies := []int{}
	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			primies = append(primies, i)
		}
	}

	return primies
}

func largestPrimeLessThan(n int, primes []int) int {
	for i := len(primes) - 1; i >= 0; i-- {
		if primes[i] <= n {
			return primes[i]
		}
	}
	return -1
}

func primeSubOperation(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	maxNum := 0

	for _, num := range nums {
		if maxNum < num {
			maxNum = num
		}
	}

	primes := sive(maxNum)

	// Traverse the array and attempt to make it strictly increasing
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			// Find the largest prime less than nums[i] and subtract it
			prime := largestPrimeLessThan(nums[i], primes)
			if prime == -1 {
				// No valid prime to subtract, return false
				return false
			}
			nums[i-1] -= prime
			// After subtraction, check if the array is still not increasing
			if nums[i] <= nums[i-1] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(primeSubOperation([]int{4, 9, 6, 10}))
	fmt.Println(primeSubOperation([]int{2, 2}))
}
