package vardic

import (
	"math"
	"strings"
)

func Min(vals ...int) int {
	min := math.MaxInt
	for _, v := range vals {
		if min > v {
			min = v
		}
	}
	return min
}

func Max(vals ...int) int {
	max := math.MinInt
	for _, v := range vals {
		if max < v {
			max = v
		}
	}
	return max
}

func Join(vals ...string) string {
	return strings.Join(vals, " ")
}
