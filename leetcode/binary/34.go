package main

func lowerBound1(nums []int, target int) int {
	left, right := -1, len(nums)
	for left+1 < right { // 开区间, 不能为空
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}

func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right { // 闭区间, 可以相等
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func searchRange(nums []int, target int) []int {
	first := lowerBound(nums, target)
	if first == len(nums) || nums[first] != target {
		return []int{-1, -1}
	}
	last := lowerBound(nums, target+1) - 1 // 由于都是整数, 最后一个位置可以看成找 target+1 的位置, 然后位置 -1 就得到了最后一个
	return []int{first, last}
}
