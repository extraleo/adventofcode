package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || q == root || p == root {
			return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
			return root
	}
	if left != nil {
			return left
	}
	return right
}
