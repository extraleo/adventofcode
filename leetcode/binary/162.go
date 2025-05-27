package main

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-2
	for left+1 < right {
		mid := (left + right) / 2
		if nums[mid] < nums[mid+1] {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}
