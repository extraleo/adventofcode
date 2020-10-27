package main

import (
	"encoding/json"
	. "extraleo/algorithm/models"
	"extraleo/algorithm/utils"
	"fmt"
	"strings"
)

// should not use tree node, you can update node
func day() int {
	data := buildMap()
	count := 0
	// for _,item:=range data{
	str, _ := json.Marshal(data["K"])
	strData, _ := json.Marshal(data)
	fmt.Printf("K \n", string(strData))
	fmt.Printf("K \n", string(str))
	fmt.Println("K", countSelf(data["K"]))
	// count +=countSelf(item)
	// }
	return count
}

func buildMap() map[string]*TreeNode {
	data := utils.SpiltInputByLine(utils.INPUT)
	treeMap := make(map[string]*TreeNode)
	for _, item := range data {
		nodeList := strings.Split(strings.Trim(item, " "), ")")
		if treeMap[nodeList[0]] == nil {
			treeMap[nodeList[0]] = &TreeNode{
				Name: nodeList[0],
			}
		}

		var child *TreeNode
		if treeMap[nodeList[1]] == nil {
			child = &TreeNode{
				Name: nodeList[1],
			}
		} else {
			child = treeMap[nodeList[1]]
		}

		if treeMap[nodeList[0]].Left == nil {
			treeMap[nodeList[0]].Left = child
		} else {
			treeMap[nodeList[0]].Right = child
		}
	}
	return treeMap
}

func countSelf(self *TreeNode) int {
	if self == nil {
		return 0
	}
	child := 0
	if self.Left != nil {
		child++
	}
	if self.Right != nil {
		child++
	}
	return child + countSelf(self.Right) + countSelf(self.Left)
}

func night() {

}

func main() {
	fmt.Println("day6: ", day())
}
