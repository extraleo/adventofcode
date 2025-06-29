package backtracking

import "slices"

func isMalache(s string, left, right int) bool{
	for left < right{
		if s [left] != s[right]{
			return false
		}
		left++
		right--
	}
	return true
}

func partition(s string) [][]string {
	ans := [][]string{}
	n := len(s)
	if n == 0 {
		return ans
	}

	path := []string{} 

	var dfs func(int)
	dfs = func(i int){
		if i == n {
			ans = append(ans, slices.Clone(path))
		}
		for right:= i ; right <n ; right++{
			if isMalache(s, i, right){
				path = append(path, s[i:right+1])
				dfs(right+1)
				path = path[:len(path)-1]

			}
		}
	}
	dfs(0)
	return ans
}