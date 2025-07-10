package hot100

import "math"

func coinChange(coins []int, amount int) int {
	mem := make([][]int, len(coins)+1)
	for i:=range(mem){
		mem[i] = make([]int, amount+1)
		for j := range(mem[i]){
			mem[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, c int) (res int){
		if i < 0 {
			if c == 0 {
				return 0
			}
			return math.MaxInt/2
		}
		p := &mem[i][c]
		if *p != -1{
			return *p
		}
        defer func(){*p =res}()	
        if c < coins[i] {
			return dfs(i-1, c)
		}
		return min(dfs(i-1, c), dfs(i, c-coins[i])+1)
	}
	ans := dfs(len(coins)-1, amount)
	if ans < math.MaxInt/2{
		return ans
	}
	return -1
}