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
	for i := range n - 2 {
		x := nums[i]
		if i > 0 && nums[i-1] == x {
			continue
		}
		// x + 最近的两个数都 > 0了, 说明这个数之后没有符合条件的数了
		if x+nums[i+1]+nums[i+2] > 0 {
			break
		}
		// x + 最远的两个数小于0了, 说明这个数不行, 换个数
		if x+nums[n-1]+nums[n-2] < 0 {
			continue
		}

		j := i + 1
		k := n - 1
		for j < k {
			s := x + nums[j] + nums[k]
			if s > 0 {
				k--
			} else if s < 0 {
				j++
			} else {
				result = append(result, []int{x, nums[j], nums[k]})
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}

	}
	return result
}
