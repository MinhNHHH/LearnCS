package main

import "fmt"

func abs(k int) int {
	if k < 0 {
		return -k
	}
	return k
}

func decrypt(code []int, k int) []int {
	res := make([]int, len(code))
	n := len(code)
	if k == 0 {
		return res
	}

	for i := 1; i <= abs(k); i++ {
		if k < 0 {
			res[0] += code[(-i+n)%n]
		} else {
			res[0] += code[(i+1)%n]
		}
	}

	for i := 1; i < len(code); i++ {
		pre := res[i-1]
		if k < 0 {
			res[i] = pre - code[(i-1+k+n)%n] + code[i-1]
		} else {
			res[i] = pre + code[(i+k)%n] - code[i]
		}
	}

	return res
}

func main() {
	fmt.Println(decrypt([]int{10, 5, 7, 7, 3, 2, 10, 3, 6, 9, 1, 6}, -4))
	// fmt.Println(decrypt([]int{5, 7, 1, 4}, 3))

}
