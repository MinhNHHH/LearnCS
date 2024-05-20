package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverse(s[:2])
	fmt.Println("frist", s)
	reverse(s[2:])
	fmt.Println("second", s)
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"

	a := make([]int, 3)[3:]
	fmt.Println(len(a), cap(a))

	var b []int
	// len(s) == 0, s == nil
	b = nil
	// len(b) == 0, b == nil
	b = []int(nil) // len(b) == 0, b == nil
	b = []int{}
	// len(b) == 0, b != nil
	fmt.Println(b)
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
