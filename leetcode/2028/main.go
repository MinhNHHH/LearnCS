package main

import "fmt"

func sumArr(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func missingRolls(rolls []int, mean int, n int) []int {
	res := make([]int, n)
	knowSum := sumArr(rolls)
	totalSum := mean * (len(rolls) + n)
	missingNum := totalSum - knowSum

	for i := 0; i < n; i++ {
		res[i] = 1
	}
	j := 0
	minNum := 1 * n
	maxNum := 6 * n
	if missingNum < minNum || missingNum > maxNum {
		return []int{}
	}
	sumRes := len(res)
	for sumRes < missingNum {
		index := j % len(res)
		res[index]++
		sumRes++
		j++
	}
	return res
}

func main() {
	fmt.Println(missingRolls([]int{1, 5, 6}, 3, 4))
}
