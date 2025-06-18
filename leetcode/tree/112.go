package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == root.Right {
		return targetSum == 0
	}
	l := hasPathSum(root.Left, targetSum)
	r := hasPathSum(root.Right, targetSum)
	return l || r
}
