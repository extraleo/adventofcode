package dp

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	mem := make([][]int, m+1)
	for i := range(mem){
			mem[i] = make([]int, n+1)
			for j := range(mem[i]){
					mem[i][j] = -1
			}
					
			
	}
	var dfs func(int, int) int
	dfs = func(i, j int)(res int) {
			if i < j {
					return 0
			}
			if j < 0 {
					return 1
			}
			p := &mem[i][j]
			if *p != -1 {
					return *p
			} 
			defer func(){*p = res}()

			res = dfs(i-1, j)
			if s[i] == t[j]{
					res += dfs(i-1, j-1)
			}
			return res
	}
	return dfs(m-1, n-1)
}