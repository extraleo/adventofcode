import "slices"

func subsets(nums []int) [][]int {
	n := len(n)
	ans := [][]int{}

	if n == 0 {
		return ans
	}

	path := []int{}

	var dfs func(int)
	dfs = func(i int){
		if i == n {
			ans = append(ans, slices.Clone(path))
			return
		}
		dfs(i+1)

		path = append(path, nums[i])
		dfs(i+1)
		path = path[:len(path)-1]
	}
	dfs(0)
	return ans
}