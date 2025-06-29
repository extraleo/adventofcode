package backtracking

func letterCombinations(digits string) []string {
    mapping := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

		path := []rune{}
		ans := []string{}
		n := len(digits)
		if n == 0 {
			return ans
		}

		var dfs func(int)
		dfs = func(i int){
			if i == n {
				ans = append(ans, string(path))
				return
			}
			for _, c := range(mapping[digits[i] - '0']){
				path[i] = c
				dfs(i+1)
			}
		}
		dfs(0)
		return ans
}
