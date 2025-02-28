package main

import (
	"adventofcode/utils"
	_ "embed"
	"fmt"
	"slices"
)

//go:embed sample-updates.txt
var updateInput string
//go:embed sample-rules.txt
var ruleInput string

func day() {
	day := 0
	updates := utils.SpiltInputGetIntList(updateInput, ",")
	rules := utils.SpiltInputGetIntList(ruleInput, "|")
	for _, u := range(updates){
		if isUpdateValid(rules, u){
			mid := len(u) / 2
			day += u[mid]
		}
	}
	fmt.Println("day: ", day)
}

func isUpdateValid(rules [][]int, update []int) bool{
	for _, r := range(rules) {
		left := r[0]
		right := r[1]
		if slices.Contains(update, left) && slices.Contains(update, right) && slices.Index(update, left) > slices.Index(update, right){
			return false
		}
	}
	return true
}

func isUpdateValidAndRule(rules [][]int, update []int)  []int{
	for _, r := range(rules) {
		left := r[0]
		right := r[1]
		lIndex := slices.Index(update, left)
		rIndex := slices.Index(update, right)
		if slices.Contains(update, left) && slices.Contains(update, right) && lIndex > rIndex{
			tmp := update[lIndex]
			update[lIndex] = update[rIndex]
			update[rIndex] = tmp
			return update
		}
	}
	return nil
}

func night(){
	night := 0
	updates := utils.SpiltInputGetIntList(updateInput, ",")
	rules := utils.SpiltInputGetIntList(ruleInput, "|")
	for _, u := range(updates){
		correct := isUpdateValidAndRule(rules, u)
		if correct != nil {
			mid := len(correct) / 2
			night += u[mid]
		}		
	}
	fmt.Println("night: ", night)
}

func main(){
	day()
	night()
}