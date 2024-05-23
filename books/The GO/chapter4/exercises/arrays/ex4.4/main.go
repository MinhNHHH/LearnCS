package main

func main() {}

func rotateRight(s []int, k int) []int {
	n := len(s)
	if n == 0 || k%n == 0 {
		return s
	}
	k = k % n // In case k is larger than the length of the array

	// Create a new slice to store the rotated elements
	rotated := make([]int, n)

	// Copy the last k elements to the beginning of the new slice
	for i := 0; i < k; i++ {
		rotated[i] = s[n-k+i]
	}

	// Copy the first n-k elements to the end of the new slice
	for i := 0; i < n-k; i++ {
		rotated[k+i] = s[i]
	}

	// Copy the rotated elements back to the original slice
	copy(s, rotated)

	return s
}
