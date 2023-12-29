package main

import (
	"adventofcode/utils/models"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	day()
	night()
}

func night() {
	panic("unimplemented")
}

func day() {
	grid := models.ParseAsMatrixFromInput(input)
	loads := make([]int, len(grid[0]))

	n := len(grid)
	for i:= range loads {
		loads[i] = n
	}

	sum := 0

	for i, r := range grid {
		for j, v := range r {
			if v == 'O' {
				sum += loads[j]
				loads[j]--
			}
			if v == '#' {
				loads[j] = n - i - 1
			}
		}
	}
	fmt.Println("day", sum)
}
