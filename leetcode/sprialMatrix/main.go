package main

import "fmt"

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	res := [][]int{}
	left, up := rStart, cStart
	right := cols - 1
	down := rows - 1
	for len(res) < rows*cols {
		for i := left; i < right; i++ {
			res = append(res, []int{left, i})
		}
		// for j := up + 1; j < down+1; j++ {
		// 	res = append(res, []int{})
		// }
		// if up != down {
		// 	for k := right - 1; k > left-1; k-- {
		// 		res = append(res, []int{down, k})
		// 	}
		// }
		// if left != right {
		// 	for l := down - 1; l > up; l-- {
		// 		res = append(res, []int{l, up})
		// 	}
		// }
		// left += 1
		// right -= 1
		// up += 1
		// down -= 1
	}

	return res
}

func main() {
	fmt.Println(spiralMatrixIII(5, 6, 1, 4))
	fmt.Println(spiralMatrixIII(1, 4, 0, 0))
}
