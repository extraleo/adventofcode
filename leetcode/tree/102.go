package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    ans := [][]int{}
    cur := []*TreeNode{root}

    for len(cur) > 0 {
         vals := []int{}
         nxt := []*TreeNode{}
         for _, node := range(cur){
            vals = append(vals, node.Val)
            if node.Left != nil {
                nxt = append(nxt, node.Left)
            }
            if node.Right != nil{
                nxt = append(nxt, node.Right)
            }
         }
         cur = nxt
         ans = append(ans, vals)
    }

    return ans
}