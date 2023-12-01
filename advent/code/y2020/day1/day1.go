package main

import (
	"adventofcode/utils"
	"fmt"
)

func day() int {
	data := utils.SpiltInputGetInt(utils.INPUT, "\n")
	dataMap := make(map[int]bool)
	for _, item := range data {
		dataMap[item] = true
	}
	for _, item := range data {
		if dataMap[2020-item] {
			return item * (2020 - item)
		}
	}
	return -1
}

func night() int {
	data := utils.SpiltInputGetInt(utils.INPUT, "\n")
	dataMap := make(map[int]bool)
	for _, item := range data {
		dataMap[item] = true
	}
	for _,item :=range data{
		second := 2020 - item
		for _, k := range data{
			if dataMap[(second-k)]{
				fmt.Print("Item: ,K, k-second:",item, k, second-k)
				return k*(second-k)*item
			}
		}
	}
	return -1
}

func main() {
	fmt.Println("Night: ", night())
}
