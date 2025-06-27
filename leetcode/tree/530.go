package main

import "math"

func getMinimumDifference(root *TreeNode) int {
	pre, ans := math.MinInt/2, math.MaxInt
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode){
			if node == nil{
					return
			}
			dfs(node.Left)
			ans = min(ans, node.Val - pre)
			pre = node.Val
			dfs(node.Right)
	}
	dfs(root)
	return ans
}