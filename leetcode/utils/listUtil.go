package utils

func SumList(nums []int, start, end int) int {
	sum := 0
	for i := start; i < end; i++ {
		sum += nums[i]
	}
	return sum
}