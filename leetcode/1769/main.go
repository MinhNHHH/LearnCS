package main

import (
	"fmt"
)

func minOperations(boxes string) []int {
	answer := make([]int, len(boxes))
	cntL, cntR, sumL, sumR := 0, 0, 0, 0

	for i := 1; i < len(boxes); i++ {
		if boxes[i] == '1' {
			cntR = cntR + 1
			sumR = sumR + i
		}
	}

	answer[0] = sumR

	for i := 1; i < len(boxes); i++ {
		sumL += cntL
		if boxes[i-1] == '1' {
			cntL = cntL + 1
			sumL = sumL + 1
		}
		if boxes[i] == '1' {
			cntR = cntR - 1
			sumR = sumR - 1
		}
		sumR -= cntR
		answer[i] = sumL + sumR
	}
	return answer
}

func main() {
	fmt.Println(minOperations("110"))
}
