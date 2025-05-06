package main

import (
	"slices"
)

func threeSum(nums []int) [][]int {
	// 顺序不重要 -> 排序先
	// i < k < j
	// 可以枚举 i, 这样从三数之和 -> 两数之和
	slices.Sort(nums)
	result := [][]int{}
	n := len(nums)
	for i := range(n-2){
		x := nums[i]
		if i > 0 && nums[i-1] == x {
			continue
		}
		j := i + 1
		k := n -1
		for ;j < k;{
			s := x + nums[j] + nums[k]
			if s == 0 {

				continue
			}
		}

	}
}
