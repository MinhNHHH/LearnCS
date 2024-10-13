package main

import "math"

type AllOne struct {
	keys     map[string]int
	maxCount int
	minCount int
	maxKey   string
	minKey   string
}

func Constructor() AllOne {
	return AllOne{
		keys:     map[string]int{},
		maxCount: math.MinInt,
		minCount: math.MaxInt,
	}
}

func (this *AllOne) Inc(key string) {
	this.keys[key]++

	if this.maxCount < this.keys[key] {
		this.maxCount = this.keys[key]
		this.maxKey = key
	}

	if this.minCount > this.keys[key] {
		this.minCount = this.keys[key]
		this.minKey = key
	}

}

func (this *AllOne) Dec(key string) {
	this.keys[key]--

	if this.maxCount < this.keys[key] {
		this.maxCount = this.keys[key]
		this.maxKey = key
	}

	if this.keys[key] == 0 {
		delete(this.keys, key)
	}

	if this.keys[key]&this.minCount > this.keys[key] {
		this.minCount = this.keys[key]
		this.minKey = key
	}
}

func (this *AllOne) GetMaxKey() string {
	return this.maxKey
}

func (this *AllOne) GetMinKey() string {
	return this.minKey
}

func main() {}
