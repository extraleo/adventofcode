package dp

func climbStairs(n int) int {
	var dfs func(int) int
	cache := make([]int, n +1)
	for x := range(n+1){
			cache[x] = -1
	}
	dfs = func(i int) int{
			if i <= 1{
					return 1
			}
			p := &cache[i] 
			if (*p == -1){
					*p =  dfs(i-1) + dfs(i-2) 
			}
			return *p
	}
	return dfs(n)
}