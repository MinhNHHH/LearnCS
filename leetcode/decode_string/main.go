// Given an encoded string, return its decoded string.

// The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

// You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc.
// Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k.
// For example, there will not be input like 3a or 2[4].

// The test cases are generated so that the length of the output will never exceed 105.

// Example 1:

// Input: s = "3[a]2[bc]"
// Output: "aaabcbc"

// Example 2:

// Input: s = "3[a2[c]]"
// Output: "accaccacc"

// Example 3:

// Input: s = "2[abc]3[cd]ef"
// Output: "abcabccdcdcdef"

// stack = [3[ab]] => times is 3
// string = ababab
// 3, "a"
package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// 91 [  | 93 ]
const s = "leetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcode"

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(len(s))
}

func decodeString(s string) string {
	stack := []rune{}
	for _, char := range s {
		curr := ""
		if char != 93 {
			stack = append(stack, char)
		} else {
			for len(stack) > 0 && stack[len(stack)-1] != 91 {
				curr = string(stack[len(stack)-1]) + curr
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			k := ""
			for len(stack) > 0 && unicode.IsNumber(stack[len(stack)-1]) {
				k = k + string(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			nums, _ := strconv.Atoi(k)
			sub := ""
			for i := 0; i < nums; i++ {
				sub = sub + curr
			}
			stack = append(stack, []rune(sub)...)
		}
	}
	return string(stack)
}
