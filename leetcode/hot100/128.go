package hot100

func longestConsecutive(nums []int) int {
	ans := 0
	has := map[int]bool{}
	for _, x := range(nums){
			has[x] = true
	}
	for k := range(has) {
			if has[k-1]{
					continue
			}
			y := k+1
			for has[y] {
					y++
			}
			ans = max(ans, y-k)
	}
	return ans
}