package main

import (
	// "fmt"
	"slices"
)

// 也是三数之和
// a + b > c && a + c > b && b + c > a
// 但是排序了 => a < b < c 这个公式就可以简化 a + b > c
// 不能简单地套公式, 就比如这题, 要枚举出两个 223
func triangleNumber(nums []int) int {
	slices.Sort(nums)
	n := len(nums)
	ans := 0
	for i := range n - 2 {
		x := nums[i]
		if x == 0 {
			continue
		}
		j := i + 1
		for k := i + 2; k < n; k++ {
			for x+nums[j] <= nums[k] {
				j++
			}
			ans += k -j
		}
	}
	return ans
}

// func main() {
// 	ans := triangleNumber([]int{2, 2, 3, 4})
// 	fmt.Println(ans)
// }
