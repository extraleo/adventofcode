package main

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	x := root.Val
	if x > high {
		return rangeSumBST(root.Left, low, high)
	}
	if x < low {
		return rangeSumBST(root.Right, low, high)
	}
	return x + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}
