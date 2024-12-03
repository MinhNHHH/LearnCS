package main

import "fmt"

func addSpaces(s string, spaces []int) string {
	newRunes := []rune{}
	i, j := 0, 0
	for _, char := range s {
		if i < len(spaces) && j == spaces[i] {
			newRunes = append(newRunes, 32)
			i++
		}
		newRunes = append(newRunes, char)
		j++

	}
	return string(newRunes)
}

func main() {
	fmt.Println(addSpaces("LeetcodeHelpsMeLearn", []int{8, 13, 15}))
}
