package main

import "fmt"

func isVaPalindrom(s string, i, j int) bool {
	if len(s) == 1 {
		return true
	}

	for i <= j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func partition(s string) [][]string {
	res := [][]string{}
	part := []string{}
	var helper func(index int)
	helper = func(index int) {
		if index >= len(s) {
			temp := []string{}
			temp = append(temp, part...)
			res = append(res, temp)
			return
		}
		for i := index; i < len(s); i++ {
			if isVaPalindrom(s, index, i) {
				part = append(part, string(s[index:i+1]))
				helper(i + 1)
				part = part[:len(part)-1]
			}
		}
	}
	helper(0)
	return res

}

func main() {
	fmt.Println(partition("abbab"))
}
