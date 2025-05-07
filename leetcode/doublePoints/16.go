package main

import (
	"math"
	"slices"
)

func threeSumClosest(nums []int, target int) int {
	slices.Sort(nums)
	n := len(nums)
	ans := 0
	min_diff := math.MaxInt
	for i := range n - 2 {
		x := nums[i]
		if i > 0 && x == nums[i-1] {
			continue
		}

		// 如果最小的三个数加起来都 > target
		sum := x + nums[i+1] + nums[i+2]
		if sum > target {
			if sum-target < min_diff {
				ans = sum
			}
			break
		}

		// 如果 x 加上最大的两个数都 < target
		sum = x + nums[n-1] + nums[n-2]
		if sum < target {
			if target-sum < min_diff {
				min_diff = target - sum
				ans = sum
			}
			continue
		}

		j, k := i+1, n-1

		for j < k {
			sum = x + nums[j] + nums[k]
			if sum == target {
				return sum
			}
			if sum > target {
				if sum-target < min_diff {
					min_diff = sum - target
					ans = sum
				}
				k--
			} else {
				if target-sum < min_diff {
					min_diff = target - sum
					ans = sum
				}
				j++
			}

		}

	}
	return ans
}
