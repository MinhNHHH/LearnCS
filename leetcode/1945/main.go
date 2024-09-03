package main

import (
	"fmt"
	"strconv"
)

func convertSToI(s string) string {
	res := ""
	for _, char := range s {
		temp := (int(char)-int('a'))%26 + 1
		iToa := strconv.Itoa(temp)
		res = res + iToa
	}
	return res
}

func getLucky(s string, k int) int {
	convert := convertSToI(s)
	res := 0
	for k > 0 {
		sum := 0
		for _, char := range convert {
			num, _ := strconv.Atoi(string(char))
			sum += num
		}
		res = sum
		convert = strconv.Itoa(sum)
		k--
	}
	return res
}

func main() {
	fmt.Println(getLucky("iiii", 1))
}
