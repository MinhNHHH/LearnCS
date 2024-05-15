package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(comma1("-12323453.1232"))
}

// 1,000
// 1,000,000.789,78

func comma1(s string) string {
	floatNum, _ := strconv.ParseFloat(s, 64)
	var buf bytes.Buffer
	if floatNum < 0 {
		buf.WriteString("-")
		s = s[1:]
	}
	splitText := strings.Split(s, ".")
	previous := comma(splitText[0])
	next := comma(splitText[1])

	buf.WriteString(previous)
	buf.WriteString(".")
	buf.WriteString(next)

	return buf.String()
}

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
