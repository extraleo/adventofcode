package main

func maxArea(height []int) int {
    ans := 0
		left, right := 0, len(height)-1
		for left < right {
				area := (right-left)*min(height[left],height[right])
				ans = max(ans, area)
				if height[left] < height[right]{
					left++
				}else{
					right--
				}
		}
		return ans
}

func min(a, b int) int{
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}