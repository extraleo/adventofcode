package main

import (
	"math"
)

// 前序遍历
// 先判断, 再递归
func isValidBST(root *TreeNode) bool {
     
	mi,_ := sufCheck(root) 
	return mi == math.MinInt
}


// 前序遍历
// 先判断, 再递归
func preCheck(root *TreeNode, left, right int) bool{
	if root == nil {
		return true
	}
	return root.Val < right && root.Val > left && preCheck(root.Left, left, root.Val) && preCheck(root.Right, root.Val, right)
}


// 中序
// left,root,right l < root < r 
func midCheck(root *TreeNode, pre int) bool {
	if root == nil {
		return true
	}
	if !midCheck(root.Left, pre){
		return false
	}
	if root.Val <= pre {
		return false
	}
	pre = root.Val
	return midCheck(root.Right,pre)
}

// 后序遍历
// check 返回这棵树的最大值和最小值,
func sufCheck(root *TreeNode) (int, int){
	if root ==  nil {
		return math.MaxInt, math.MinInt
	}
	lMin,lMax := sufCheck(root.Left)
	rMin,rMax := sufCheck(root.Right)
	v := root.Val
	if v >= rMin || v <= lMax {
		return math.MinInt,math.MaxInt
	}
	return min(lMin,v),max(rMax,v)


}