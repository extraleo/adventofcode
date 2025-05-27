package main


func numSubarrayProductLessThanK(nums []int, k int) int {
    if k <= 1{
        return 0
    }
    left, prod := 0, 1
    ans := 0
    for i, v := range(nums){
        prod *= v
        for prod >= k {
            prod /= nums[left]
            left++
        }
        ans += i - left +1
    }
    return ans
}