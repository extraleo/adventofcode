package main

import (
	"slices"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	slices.Sort(potions)
	ans := make([]int, len(spells))
	for i, v := range(spells) {
		target := int(success-1)/v +1
		ans[i]=findLargerBound(potions, target)
	}
	return ans
}

func findLargerBound(nums []int, target int) int{
	left, right := 0, len(nums) -1
	for left<=right{
		mid := (left+right)/2
		if nums[mid] < target{
			left = mid +1
		}else{
			right = mid -1
		}
	}
	if left <= 0 {
		return len(nums)
	}
	if left >= len(nums){
		return 0
	}

	return len(nums) - left
}