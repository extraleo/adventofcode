package dp

// dp[i] = max(dp[i-1], dp[i-2]+num[i])
func rob(nums []int) int {
    f0, f1 := 0, 0
    for _, x := range(nums){
      newF := max(f1, f0+x)
			f0=f1
			f1=newF
    }
    return f1
}
