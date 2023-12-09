package main

import (
	"adventofcode/utils"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func day() {
	result := 0
	lines := utils.SpiltInputGetIntList(input, ` `)
	for _, line := range lines {
		result += findExtrapolated(line)
	}
	fmt.Println("day:", result)
}

func findExtrapolated(history []int) int {
	if history[0] == history[len(history)-1] {
		return history[len(history)-1]
	}
	next := getNext(history)
	return findExtrapolated(next) + history[len(history)-1]

}

func findPreExtrapolated(history []int) int {
	if history[0] == history[len(history)-1] {
		return history[len(history)-1]
	}
	next := getNext(history)
	return history[0] - findPreExtrapolated(next)

}

func getNext(history []int) []int {
	result := []int{}
	for i := 1; i < len(history); i++ {
		result = append(result, history[i]-history[i-1])
	}
	return result
}

func night() {
	result := 0
	lines := utils.SpiltInputGetIntList(input, ` `)
	for _, line := range lines {
		result += findPreExtrapolated(line)
	}
	fmt.Println("night:", result)
}

func main() {
	night()
}
