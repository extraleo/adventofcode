package main


func trap(height []int) int {
	n := len(height)
	left, right := 0, n-1
	preMax, sufMax := 0, 0
	ans := 0
	for left <= right {
		preMax = max(height[left], preMax)
		sufMax = max(height[right], sufMax)
		if preMax < sufMax {
			ans += preMax - height[left]
			left++
		}else{
			ans += sufMax - height[right] 
			right--
		}
	}
	return ans

}

// 单调栈
func trapStack(height []int) int {
	ans := 0
	st := []int{}
	for i, h := range(height){
			for len(st) > 0 && h >= height[st[len(st)-1]]{
					bottomH := height[st[len(st)-1]] // 需要找到左边的柱子
					st = st[:len(st)-1]
					if len(st) == 0 {
							break
					}
					left := st[len(st)-1]
					dh := min(height[left], h) - bottomH
					ans += dh*(i-left-1)
			}
			st = append(st, i)
	}
	return ans
}


// 第一版
// 两次 for 循环
func trapFirst(height []int) int {
	n := len(height)
	ans := 0
	preMax := make([]int, n)
	sufMax := make([]int, n)

	preMax[0] = height[0]
	for i := 1; i < n; i++ {
		preMax[i] = max(preMax[i-1], height[i])
		
	}
	sufMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		sufMax[i] = max(sufMax[i+1], height[i])
	}

	for i := 0; i < n; i++ {
		ans += min(preMax[i], sufMax[i]) - height[i]
	}
	return ans
}
