package main

// 1 * 2 * 3 * 4 = 24
// 24 % 1 * 1
// 24 % 2
//	24 % 6
//
//
func productExceptSelf(nums []int) []int {
	pre, post := 1, 1
	res := make([]int, len(nums))

	for index, value := range nums {
		res[index] = pre
		pre = pre * value
	}

	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = post * res[i]
		post = post * nums[i]
	}
	return res
}

func main() {
	productExceptSelf([]int{1, 2, 3, 4})
}
