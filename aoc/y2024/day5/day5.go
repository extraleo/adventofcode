package main

import (
	"adventofcode/utils"
	_ "embed"
	"fmt"
	"slices"

	"github.com/dominikbraun/graph"
)

//go:embed input-updates.txt
var updateInput string
//go:embed input-rules.txt
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


func night(){
	night := 0
	updates := utils.SpiltInputGetIntList(updateInput, ",")
	rules := utils.SpiltInputGetIntList(ruleInput, "|")
	for _, u := range(updates){
		if !isUpdateValid(rules,u) {
			sort:=graph.New(graph.IntHash, graph.Directed(), graph.PreventCycles())
			for _, rule := range(rules){
				l:=rule[0]
				r:=rule[1]
				if slices.Contains(u, l) && slices.Contains(u, r){
					sort.AddVertex(l)
					sort.AddVertex(r)
					sort.AddEdge(l,r)
				}
			}
			order,_:=graph.TopologicalSort(sort)
			mid:=len(order)/2
			night += order[mid]
		}
	
	}
	fmt.Println("night: ", night)
}

func main(){
	day()
	night()
}