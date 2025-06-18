package main

func sumNumbers(root *TreeNode) int {
	return sumNums(root, 0)
}

func sumNums(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	sum = root.Val + sum*10
	if root.Left == root.Right {
		return sum
	}
	return sumNums(root.Left, sum) + sumNums(root.Right, sum)
}
