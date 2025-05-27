package main

func maximumCount(nums []int) int {
	zero := findLowerBound(nums, 0)
	if zero >= len(nums) || zero < 0 {
		return len(nums)
	}
	one := findLowerBound(nums, 1)
	if len(nums)-one > zero {
		return len(nums) - one
	}
	return zero
}

func findLowerBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
