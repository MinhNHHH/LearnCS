package main

import (
	"fmt"
)

func findDifferentBinaryString(nums []string) string {
	n := len(nums[0])
	hashMap := map[string]bool{}
	for _, num := range nums {
		hashMap[num] = true
	}
	temp := []string{}
	var helper func(curr string)
	helper = func(curr string) {
		if len(curr) == n {
			temp = append(temp, curr)
			return
		}

		helper(curr + string("0"))
		helper(curr + string("1"))
	}
	helper("")

	for _, str := range temp {
		if !hashMap[str] {
			return str
		}
	}
	return ""
}

func main() {
	fmt.Println(findDifferentBinaryString([]string{"01", "10"}))
}
