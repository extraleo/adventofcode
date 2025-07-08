package main

import "slices"

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}
	cur := []*TreeNode{root}
	for len(cur) != 0 {
		next := []*TreeNode{}
		vals := []int{}
		for _, node := range cur {
			vals = append(vals, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}

		}
		cur = next
		if len(ans)%2 != 0 {
			slices.Reverse(vals)
		}
		ans = append(ans, vals)
	}
	return ans
}
