package main

import "fmt"

func getHappyString(n int, k int) string {
	res := []string{}
	var helper func(curr string)
	helper = func(curr string) {
		if len(curr) == n {
			res = append(res, curr)
			return
		}

		for _, char := range "abc" {
			if curr == "" || curr[len(curr)-1] != byte(char) {
				helper(curr + string(char))
			}
		}
	}
	helper("")

	if len(res) >= k {
		return res[k-1]
	}
	return ""
}

func main() {
	fmt.Println(getHappyString(3, 9))
}
