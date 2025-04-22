package main

import (
	"adventofcode/leetcode/utils"
)

func findMaxAverage(nums []int, k int) float64 {
	if k >= len(nums) {
		return float64(sumList(nums, 0, len(nums))) / float64(len(nums))
	}
	based := sumList(nums, 0, k)
	max := based
	for i := k; i < len(nums); i++ {
		based = based + nums[i] - nums[i-k]
		max = utils.MaxInt(max, based)
	}
	return float64(max) / float64(k)
}

func sumList(nums []int, start, end int) int {
	sum := 0
	for i := start; i < end; i++ {
		sum += nums[i]
	}
	return sum
}

func numOfSubarrays(arr []int, k int, threshold int) int {
    thr := threshold*k
		if k > len(arr) {
			if sumList(arr, 0 , len(arr)) >= thr {
				return 1
			}else {
				return 0
			}
		}
		based := sumList(arr, 0, k)
		result := 0
		if based >= thr {
			result++
		}
		for i:=k; i< len(arr); i++{

			based = based + arr[i] - arr[i-k]
						if (based >= thr){
				result++
			}
		}
		return result
		
}
