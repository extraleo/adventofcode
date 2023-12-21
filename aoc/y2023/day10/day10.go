package main

import (
	"adventofcode/utils/models"
	"adventofcode/utils/set"
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

func main(){
	day()
	night()
}

const (
	EMPTY       = '.'
	UPPER_LEFT  = 'F'
	UPPER_RIGHT = '7'
	LOWER_LEFT  = 'L'
	LOWER_RIGHT = 'J'
	START       = 'S'
	VERTICAL    = '|'
	HORIZONTAL  = '-'
	NORTH       = iota
	SOUTH
	EAST
	WEST
)



func buildGrid(lines []string) models.Grid{
	grid:=make(Grid)
	for j,l:=range lines{
		for i, c:=range l{
			grid[Point{X: i, Y: j}] = int(c)
		}
	}
	return grid
}

func findStart(grid Grid) Point{
	for point,v:=range grid{
		if v == START{
			return point
		}
	}
	panic("cannot happen")
}


func findCycle(grid Grid, neighbor Point, direction int)(path set.Set[Point], bool){

}

func night() {
	panic("unimplemented")
}

func day() {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	grid := buildGrid(lines)
	start := findStart(grid)

	neighbors := []Point{
		{start.X + 1, start.Y},
		{start.X - 1, start.Y},
		{start.X, start.Y - 1},
		{start.X, start.Y + 1},
	}

	around := []int{WEST, EAST, SOUTH, NORTH}

	for i, n := range neighbors {
		loop, found := findCycle(grid, n, around[i])
		if found {
			l:= len(loop)
			if l%2 == 0 {
				return l / 2
			}
			return l/2 + 1
		}
	}
	panic("no solution found")


}