package main

import "sort"

type MyCalendar struct {
	calendars [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{
		calendars: [][2]int{},
	}
}

func (this *MyCalendar) Sort() {
	sort.Slice(this.calendars, func(i, j int) bool {
		return this.calendars[i][0] < this.calendars[j][0]
	})
}

func (this *MyCalendar) Book(start int, end int) bool {
	if len(this.calendars) > 0 {
		for _, calendar := range this.calendars {
			if calendar[0] < end && start < calendar[1] {
				return false
			}
		}
	}
	this.calendars = append(this.calendars, [2]int{start, end})
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */

func main() {}
