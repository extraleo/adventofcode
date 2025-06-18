package main

import "slices"

func divideArray(nums []int, k int) [][]int {
	slices.Sort(nums)
	ans := [][]int{}
	for a := range slices.Chunk(nums, 3) {
		if a[2]-a[0] > k {
			return nil

		}
		ans = append(ans, a)
	}
	return ans
}
