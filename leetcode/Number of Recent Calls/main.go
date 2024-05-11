package main

import "fmt"

func main() {
	obj := Constructor()
	param_1 := obj.Ping(1)
	param_2 := obj.Ping(100)
	param_3 := obj.Ping(3001)
	param_5 := obj.Ping(4002)
	fmt.Println(param_1, param_2, param_3, param_5)
}

type RecentCounter struct {
	requests []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

// [[-2999,1], [-2900,100], [1,3001], [2,3002]]
// [1, ,100, 3001, 3002]
func (this *RecentCounter) Ping(t int) int {
	this.requests = append(this.requests, t)
	rangeTime := []int{t - 3000, t}
	count := 0
	fmt.Println(rangeTime)
	for _, rq := range this.requests {
		if rq >= rangeTime[0] && rq <= rangeTime[1] {
			count++
		}
	}
	return count
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
