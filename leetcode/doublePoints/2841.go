package main

import (
	"adventofcode/leetcode/utils"
)

func maxSum(nums []int, m int, k int) int64 {
	memoryMap := make(map[int]int, len(nums))
	sum := 0
	max := 0
	for i, v := range nums {
		sum += v
		if replicas, ok := memoryMap[v]; ok {
			memoryMap[v] = replicas + 1
		} else {
			memoryMap[v] = 1
		}
		if i < k-1 {
			continue
		}
		max = uniqueArray(memoryMap, k, m, max, sum)
		sum -= nums[i-k+1]
		memoryMap[nums[i-k+1]] = memoryMap[nums[i-k+1]] - 1
		if memoryMap[nums[i-k+1]] == 0 {
			delete(memoryMap, nums[i-k+1])
		}

	}
	return int64(max)
}

func uniqueArray(mem map[int]int, k, m, max, sum int) int {
	if len(mem) >= m {
		return utils.MaxInt(max, sum)
	}
	return max
}
