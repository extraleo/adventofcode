package hot100

func change(amount int, coins []int) int {
	cache := make([][]int, len(coins))
	for i := range cache{
			cache[i] = make([]int, amount+1)
			for j := range(cache[i]){
					cache[i][j] = -1
			}
	}
	var dfs func(int ,int) int 
	dfs = func(i, c int) (res int) {
			if i < 0 {
					if c == 0 {
							return 1
					}
					return 0
			}

			p := &cache[i][c]
			if *p != -1 {
					return *p
			} 
			defer func(){*p=res}()

			if c < coins[i] {
					return dfs(i-1, c)
			}
			return dfs(i-1, c) + dfs(i, c-coins[i])
	}
	return dfs(len(coins)-1,amount)
}







