package main

import (
	"adventofcode/utils/models"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func buildGrid(lines []string, expand bool) (grid models.Grid, offset [][]int) {
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

	offset = append(offset, upX, upY)
	if expand {
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
	}

	count := 1
	for j, l := range lines {
		for i, c := range l {
			if c == '#' {
				grid[models.Position{X: i, Y: j}] = count
				count++
			}

		}
	}
	return
}

func day() {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	grid, _ := buildGrid(lines, true)

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

func night() {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	grid, offset := buildGrid(lines, false)
	offsetX := offset[0]
	offsetY := offset[1]
	steps := 0
	hasCalc := make(map[int]bool)
	for k, v := range grid {
		for ik, iv := range grid {
			if hasCalc[v] || hasCalc[iv] || iv == v {
			} else {
				xStart, xEnd, yStart, yEnd := 0, 0, 0, 0
				if k.X > ik.X {
					xStart, xEnd = ik.X, k.X
				} else {
					xStart, xEnd = k.X, ik.X
				}

				if k.Y > ik.Y {
					yStart, yEnd = ik.Y, k.Y
				} else {
					yStart, yEnd = k.Y, ik.Y
				}

				offsX := 0
				for i := xStart; i < xEnd; i++ {
					if offsetX[i] == 0 {
						offsX++
					}
				}
				offsY := 0
				for y := yStart; y < yEnd; y++ {
					if offsetY[y] == 1 {
						offsY++
					}
				}

				steps += xEnd + offsX*(1000000-1) - xStart + yEnd + offsY*(1000000-1) - yStart

			}
		}
		hasCalc[v] = true
	}

	fmt.Println("night:", steps)
}

func main() {
	// day()
	night()
}
