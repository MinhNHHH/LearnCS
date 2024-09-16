package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func convertStringTimeToInt(time string) int {
	timeSplit := strings.Split(time, ":")
	hour, _ := strconv.Atoi(timeSplit[0])
	minute, _ := strconv.Atoi(timeSplit[1])
	return hour*60 + minute
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMinDifference(timePoints []string) int {
	minute := []int{}
	for _, timePoint := range timePoints {
		minute = append(minute, convertStringTimeToInt(timePoint))
	}
	sort.Ints(minute)

	res := math.MaxInt
	if minute[0] >= 0 && minute[0] <= 720 {
		res = 1440 + minute[0] - minute[len(minute)-1]
	}

	for i := 1; i < len(minute); i++ {
		res = min(res, minute[i]-minute[i-1])
	}
	return res
}

func main() {
	fmt.Println(findMinDifference([]string{"23:59", "00:00", "01:00", "02:00"}))
	fmt.Println(findMinDifference([]string{"05:31", "21:08", "01:35"}))
}
