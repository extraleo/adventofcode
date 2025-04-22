package main

import (
	"adventofcode/leetcode/utils"
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
	if k >= len(nums){
		return 	 float64(sumList(nums, 0 ,len(nums)))/float64(len(nums))

	}
	based := sumList(nums, 0, k)
	max := based
	for i:=k; i < len(nums); i++ {
		based = based +nums[i]-nums[i-k]
		max = utils.MaxInt(max, based)
	}
	return float64(max)/float64(k)
}

func sumList(nums []int, start, end int) int{
	sum := 0
	for i:=start; i< end; i ++{
		sum += nums[i]
	}
	return sum
}


func main(){
	fmt.Println(findMaxAverage([]int{0,4,0,3,2}, 1))
}