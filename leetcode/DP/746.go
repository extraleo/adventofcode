package dp

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	var dfs func(int) int
	cache := make([]int, n+1)
	for x := range(n+1){
			cache[x] = -1
	}
	dfs = func(i int) int {
			if i<= 1 {
					return 0
			}
			p := &cache[i]
			if *p == -1 {
					*p = min(dfs(i-1)+cost[i-1], dfs(i-2)+cost[i-2])
			}
			return *p
	}
	return dfs(n)
}