package hot100

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	reverseList(nums, 0, n-1)
	reverseList(nums, 0, k-1)
	reverseList(nums, k, n-1)

}

func reverseList(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start += 1
		end -= 1
	}
}
