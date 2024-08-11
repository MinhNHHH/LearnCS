package main

import (
	"fmt"
)

func numberToWords(num int) string {
	tempNumber := []int{1000000000, 1000000, 1000, 100, 90, 80, 70, 60, 50, 40, 30, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	tempWords := []string{"Billion", "Million", "Thousand", "Hundred", "Ninety", "Eighty", "Seventy", "Sixty", "Fifty", "Fourty", "Thirty", "Twenty", "Nineteen", "Eighteen", "Seventeen", "Sixteen", "Fifteen", "Fourteen", "Thirteen", "Twelve", "Eleven", "Ten", "Night", "Eight", "Seven", "Six", "Five", "Four", "Three", "Two", "One", "Zero"}
	words := ""
	i := 0
	for num > 0 {
		for num >= tempNumber[i] {
			words += tempWords[i]
			num -= tempNumber[i]
		}
		i++
	}
	return words
}

func main() {
	fmt.Println(numberToWords(123))
	fmt.Println(numberToWords(12345))
}
