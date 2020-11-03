package structures

type TreeNode struct {
	Val int
	Name string `json: name`
	Left *TreeNode `json: left`
	Right *TreeNode `json: right`
}

