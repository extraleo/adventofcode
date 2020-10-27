package main

import (
	. "extraleo/algorithm/models"
	"extraleo/algorithm/utils"
	"fmt"
	"strings"
)

func day() int {
	data := utils.SpiltInputByLine(utils.INPUT)
	treeMap := make(map[string]*TreeNode)
	for _, item := range data {
		nodeList := strings.Split(strings.Trim(item," "), ")")

		if treeMap[nodeList[0]] == nil {
			treeMap[nodeList[0]] = &TreeNode{
				Name: nodeList[0],
			}
		}
		if treeMap[nodeList[1]] == nil {
			treeMap[nodeList[1]] = &TreeNode{
				Val: treeMap[nodeList[0]].Val + 1,
				Name: nodeList[1],
			}
		}else {
			treeMap[nodeList[1]].Val = treeMap[nodeList[1]].Val + 1
		}

		if treeMap[nodeList[0]].Left == nil {
			treeMap[nodeList[0]].Left = treeMap[nodeList[1]]
		} else {
			treeMap[nodeList[0]].Right = treeMap[nodeList[1]]
		}
	}

	count := 0
	for _,value:=range treeMap{
		count = count + value.Val
		fmt.Println(value.Name, value.Val)
	}

	return count
}

func night() {

}

func main() {
	fmt.Println("day6: ", day())
}
