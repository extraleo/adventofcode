package main

func isSymmetric(root *TreeNode) bool {
	return isSame(root.Left, root.Right)
}

func isSame(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSame(p.Left, q.Right) && isSame(p.Right, q.Left)
}
