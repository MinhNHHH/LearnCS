package main

import "fmt"

func getMax(arr []int) (int, int) {
	max := arr[0]
	i := 0
	for index, value := range arr {
		if max < value {
			max = value
			i = index
		}
	}
	return max, i
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func maxPoints(points [][]int) int64 {
	current := make([]int64, len(points[0]))
	pervious := make([]int64, len(points[0]))
	var maxScore int64

	peak := 0
	for _, gird := range points {
		fmt.Println(gird)
		for i := 0; i < len(points[0]); i++ {
			peak = max(peak-1, pervious[i])
			current[i] = peak
		}
		peak = 0
		for j := 0; j < len(points[0]); j++ {
			peak = max(peak-1, pervious[j])
			if current[j] > peak {
				current[j] += int64(gird[j])
			} else {
				current[i] = peak + int64(level[i])
			}
		}
		copy(pervious, current)
		fmt.Print(current)
	}
	return maxScore
}

func main() {
	fmt.Println(maxPoints([][]int{{1, 2, 3}, {1, 5, 1}, {3, 1, 1}}))
	fmt.Println(maxPoints([][]int{{0, 3, 0, 4, 2}, {5, 4, 2, 4, 1}, {5, 0, 0, 5, 1}, {2, 0, 1, 0, 3}}))
}
