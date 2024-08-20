package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	r := 0
	res := ""
	for r < numRows {
		increment := 2 * (numRows - 1)
		for i := r; i < len(s); i += increment {
			res = res + string(s[i])
			if r > 0 && r < numRows && i+increment-2*r < len(s) {
				res = res + string(s[i+increment-2*r])
			}
		}
		r++
		strings.ToLower()
	}
	return res
}

func main() {
	// fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 3))
}
