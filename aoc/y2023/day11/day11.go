package main

import (
	"adventofcode/utils/models"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func buildGrid(lines []string) (grid models.Grid) {
	grid = make(models.Grid)
	upX := make([]int, len(lines[0]))
	upY := make([]int, len(lines))
	for j, l := range lines {
		var noGalaxy = true
		for i, c := range l {
			if c == '#' {
				noGalaxy = false
				upX[i] = 1
			}
			if noGalaxy && i == len(l)-1 {
				upY[j] = 1
			}
		}
	}

	offsetY := 0
	for y := range upY {
		if upY[y] == 1 {
			empty := lines[y+offsetY]
			lines = append(lines[:y+offsetY+1], lines[y+offsetY:]...)
			lines[y+offsetY+1] = empty
			offsetY++
		}
	}

	offsetX := 0
	for x := range upX {
		if upX[x] == 0 {
			for i, line := range lines {
				index := x + offsetX
				newLine := line[:index] + "." + line[index:]
				lines[i] = newLine
			}
			offsetX++

		}
	}
	count := 1
	for j, l := range lines {
		for i, c := range l {
			if c == '#' {
				grid[models.Point{X: i, Y: j}] = count
				count++
			}

		}
	}
	return
}

func day() {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	grid := buildGrid(lines)

	steps := 0
	hasCalc := make(map[int]bool)
	for k, v := range grid {
		for ik, iv := range grid {
			if hasCalc[v] || hasCalc[iv] {

			} else {
				steps += k.Manhattan(ik)

			}
		}
		hasCalc[v] = true
	}

	fmt.Println("day:", steps)

}

func main() {
	day()
}
