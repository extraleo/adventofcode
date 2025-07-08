package main


 func findBottomLeftValue(root *TreeNode) int {
	path := _levelOrder(root)
	last := path[len(path)-1]
	return last[0]
}

func _levelOrder(root *TreeNode) [][]int {
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
	ans = append(ans, vals)
}
return ans
}