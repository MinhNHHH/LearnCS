package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("123456789"))
}

// 1,000
// 1,000,000
func comma(s string) string {
	var buf bytes.Buffer
	count := 0
	var result string
	for i := len(s) - 1; i >= 0; i-- {
		if count == 3 {
			result = "," + result
			count = 0
		}
		result = string(s[i]) + result
		count++
	}
	buf.WriteString(result)
	return buf.String()
}
