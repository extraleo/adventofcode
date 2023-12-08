package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func day() {

	lines := utils.SpiltInputByLine("input.txt")
	times := utils.ConvertIntString2IntList(strings.Split(lines[0], ":")[1])
	distances := utils.ConvertIntString2IntList(strings.Split(lines[1], ":")[1])
	result := 1
	for round, time := range times {
		count := 0
		for hold := 0; hold <= time; hold++ {
			if race(hold, time) > distances[round] {
				count++
			}
		}
		if (count>0){
			result = result * count
		}
	}
	fmt.Println("day:", result)
}

func 

func race(hold, total int) int {
	return hold * (total - hold)
}

func main() {
	day()
}
