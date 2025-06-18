package main

import "math"

func goodNodes(root *TreeNode) int {
	return goodNode(root, math.MinInt)
}

func goodNode(root *TreeNode, maxNum int) int {
	if root == nil {
		return 0
	}
	maxNum = max(root.Val, maxNum)
	if root.Val >= maxNum {
		return goodNode(root.Left, maxNum) + goodNode(root.Right, maxNum) + 1
	}
	return goodNode(root.Left, maxNum) + goodNode(root.Right, maxNum)
}
