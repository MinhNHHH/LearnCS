package main

import "fmt"

type MyCalendarTwo struct {
	singleBooks [][2]int
	doubleBooks [][2]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		singleBooks: [][2]int{},
		doubleBooks: [][2]int{},
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, doubleBook := range this.doubleBooks {
		start1, end1 := doubleBook[0], doubleBook[1]
		if start1 < end && end1 > start {
			return false
		}
	}

	for _, singleBook := range this.singleBooks {
		start1, end1 := singleBook[0], singleBook[1]
		if start1 < end && end1 > start {
			this.doubleBooks = append(this.doubleBooks, [2]int{max(start1, start), min(end1, end)})
		}
	}
	this.singleBooks = append(this.singleBooks, [2]int{start, end})
	return true
}

func main() {
	mycalen := Constructor()
	listBooks := [][]int{{10, 20}, {50, 60}, {10, 40}, {5, 15}, {5, 10}, {25, 55}}

	for _, book := range listBooks {
		fmt.Println(mycalen.Book(book[0], book[1]))
	}
	fmt.Println("singleBook", mycalen.singleBooks)
	fmt.Println("doubleBooks", mycalen.doubleBooks)
}
