package main

import "fmt"

func maxDivision(a, b int) int {
	for a%b == 0 {
		a = a / b
	}
	return a
}

func isUgly(n int) int {
	check := maxDivision(n, 2)
	check = maxDivision(check, 3)
	check = maxDivision(check, 5)
	return check
}

// Native Solution O(nLog(n))
//func nthUglyNumber(n int) int {
//	res := 1
//	i := 0
//	number := 1
//	for i < n {
//		if isUgly(number) == 1 {
//			i++
//			res = number
//		}
//		number++
//	}
//	return res
//}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}

func nthUglyNumber(n int) int {
	uglynumbers := make([]int, n)
	uglynumbers[0] = 1
	p2, p3, p5 := 0, 0, 0

	for i := 1; i < n; i++ {
		nump2 := uglynumbers[p2] * 2
		nump3 := uglynumbers[p3] * 3
		nump5 := uglynumbers[p5] * 5
		nextugly := min(nump2, nump3, nump5)
		uglynumbers[i] = nextugly

		if nextugly == nump2 {
			p2++
		}
		if nextugly == nump3 {
			p3++
		}
		if nextugly == nump5 {
			p5++
		}
	}
	return uglynumbers[n-1]
}

func main() {
	fmt.Println(nthUglyNumber(11))
}
