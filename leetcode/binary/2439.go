package main


func minimizeArrayValue(nums []int) int {
	sum,ans := 0, 0
	
	for i,v := range(nums)	{
		sum += v
		ans = max(ans, (sum+i)/(i+1))
	}
	return ans
}

