package main

func getAverages(nums []int, k int) []int {
	result := make([]int, 0)
	based := 0
	for i:=0; i< len(nums); i++{
		result = append(result, -1)
		based += nums[i]
		if i - 2*k > 0 {
			based -= nums[i-(2*k+1)]
		}
		if i -2*k >=0 {
			result[i-k]= based/(2*k+1)
		}
	}
		return result
}