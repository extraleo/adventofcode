package main

func closestNodes(root *TreeNode, queries []int) [][]int {

	// 中序遍历
	nodes := []int{}
	var dfs func(*TreeNode)
	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		dfs(tn.Left)
		nodes = append(nodes, tn.Val)
		dfs(tn.Right)
	}
	dfs(root)

	// 二分查找
	n := len(nodes)
	ans := make([][]int, len(queries))

	for i,q := range(queries){
		left,right := 0,n-1
		for left <= right {
			mid := (left+right)/2
			if nodes[mid] > q {
				left = mid + 1
			}else{
				right = mid -1
			}
		}

		if left < 0 {
			ans[i] = []int{-1,nodes[0]}
		}else if left > n {
			ans[i] = []int{nodes[n-1], -1}
		}else{
			ans[i] = []int{nodes[left-1],nodes[left]}
		}
	}
	return ans
}
