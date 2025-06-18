package main

func isBalanced(root *TreeNode) bool {
	return height(root) != -1

}

func height(root *TreeNode) int{
	if root == nil {
			return 0
	}
	l := height(root.Left)
	if l == -1 {
			return -1
	}
	r := height(root.Right)
	if r == -1 || abs(l-r) > 1{
			return -1
	}
	return max(r, l) +1
}

func abs(a int)int{
	if a >0 {
			return  a
	}
	return -a
}