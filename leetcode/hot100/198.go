package hot100

func rob(nums []int) int {
	mem := make([]int, len(nums)+1)
	for i := range(mem){
		mem[i] = -1
	}
    var dfs func(int) int
		dfs = func(i int) (res int){
			if i < 0 {
				return 0 
			}
			p := &mem[i]
			if *p != -1{
				return *p
			}
			defer func(){*p = res}()
			return max(dfs(i-2) + nums[i], dfs(i-1))
			
		}
		return dfs(len(nums)-1)
}