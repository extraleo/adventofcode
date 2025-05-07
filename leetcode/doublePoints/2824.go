package main

import "slices"

func countPairs(nums []int, target int) int {
	slices.Sort(nums)
	n := len(nums)
	count := 0
	if nums[n-2]+nums[n-1] < target {
		return n * (n - 1) / 2
	}

	for i := 0; i < n; i++ {
		if nums[i]+nums[i+1] >= target {
			break
		}
		r := n - 1
		for i < r {
			x := nums[i] + nums[r]
			if x < target {
				count += (r - i)
				break
			} else {
				r--
			}
		}

	}
	return count
}

// func main() {
// 	countPairs([]int{-1, 1, 2, 3, 1}, 2)
// }
