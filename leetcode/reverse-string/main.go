package main

import (
	"fmt"
	"strings"
)

func removeSpaces(input string) []string {
	splitText := strings.Split(input, " ")
	textRemvedSpaces := []string{}
	for i := len(splitText) - 1; i >= 0; i-- {
		if splitText[i] != "" {
			textRemvedSpaces = append(textRemvedSpaces, splitText[i])
		}
	}
	return textRemvedSpaces
}

func reverseWords(s string) string {
	textRemovedSpaces := removeSpaces(s)
	res := ""
	for _, str := range textRemovedSpaces {
		res = res + str + " "
	}
	return strings.Trim(res, " ")
}

func main() {
	fmt.Println(strings.Trim("sa das    ", " "), "adasd")
}
