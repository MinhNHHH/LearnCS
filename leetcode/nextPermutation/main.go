package main

func reverse(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

func nextPermutation(nums []int) {
	var pivot int

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > nums[i-1] {
			pivot = i - 1
			break
		} else {
			nums = reverse(nums)
			return
		}
	}

	swap := len(nums) - 1
	for nums[swap] <= nums[pivot] {
		swap -= 1
	}
	nums[pivot], nums[swap] = nums[swap], nums[pivot]

	nums = append(nums[:pivot+1], reverse(nums[pivot+1:])...)
}

func main() {
	nextPermutation([]int{1, 2, 3})
}
