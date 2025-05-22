package main

func minSubArrayLen(target int, nums []int) int {
	left, sum := 0, 0
	ans := len(nums) + 1
	for i, v := range nums {
		sum += v
		for sum >= target {
			ans = min(ans, i-left+1)
			sum -= nums[left]
			left++

		}
	}
	if ans <= len(nums) {
		return ans
	}
	return 0
}
