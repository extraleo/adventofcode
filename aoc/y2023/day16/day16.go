package main

import (
	"adventofcode/utils/models"
	"adventofcode/utils/set"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
	UPDOWN
	LEFTRIGHT
)

type state struct{
	dir int
	local models.Position
}


func main() {
	day()
	night()
}

func night() {
	grid := models.ParseAsMatrixFromInput(input)
	var night int
	for x := 0; x < len(grid[0]); x++ {
		night = max(night, find(state{local: models.Position{X: x, Y: 0}, dir: DOWN}, grid))
		night = max(night, find(state{local: models.Position{X: x, Y: len(grid)-1}, dir: UP}, grid))
	}
	for y := 0; y < len(grid); y++ {
		night = max(night, find(state{local: models.Position{X: 0, Y: y}, dir: RIGHT}, grid))
		night = max(night, find(state{local: models.Position{X: len(grid[0])-1, Y: y}, dir: LEFT}, grid))
	}
	fmt.Println("night", night)
}

func day() {
	grid := models.ParseAsMatrixFromInput(input)
	start := state{local: models.Position{X: 0, Y: 0}, dir: RIGHT}

	day := find(start, grid)
	fmt.Println("day", day)

}

func find(start state, grid models.Matrix) int {
	queue := []state{start}
	visited := set.NewSet[state]()
	energized := set.NewSet[models.Position]()
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if visited.Contains(current) || !grid.IsValidPosition(current.local) {
			continue
		}
		visited.Add(current)
		energized.Add(current.local)
		cx, cy := current.local.X, current.local.Y
		switch nextDirection(current.dir, grid[cy][cx]) {
		case UP:
			queue = append(queue, state{local: models.Position{X: cx, Y: cy - 1}, dir: UP})
		case RIGHT:
			queue = append(queue, state{local: models.Position{X: cx + 1, Y: cy}, dir: RIGHT})
		case DOWN:
			queue = append(queue, state{local: models.Position{X: cx, Y: cy + 1}, dir: DOWN})
		case LEFT:
			queue = append(queue, state{local: models.Position{X: cx - 1, Y: cy}, dir: LEFT})
		case UPDOWN:
			queue = append(queue, state{local: models.Position{X: cx, Y: cy - 1}, dir: UP})
			queue = append(queue, state{local: models.Position{X: cx, Y: cy + 1}, dir: DOWN})
		case LEFTRIGHT:
			queue = append(queue, state{local: models.Position{X: cx - 1, Y: cy}, dir: LEFT})
			queue = append(queue, state{local: models.Position{X: cx + 1, Y: cy}, dir: RIGHT})
		}
	}
	return energized.Len()
}



func nextDirection(currentDir int, nextChar rune) int{
	switch currentDir {
	case RIGHT:
		switch nextChar {
		case '.', '-':
			return RIGHT
		case '|':
			return UPDOWN
		case '/':
			return UP
		case '\\':
			return DOWN
		}
	case LEFT:
		switch nextChar {
		case '.', '-':
			return LEFT
		case '|':
			return UPDOWN
		case '/':
			return DOWN
		case '\\':
			return UP
		}
	case UP:
		switch nextChar {
		case '.', '|':
			return UP
		case '-':
			return LEFTRIGHT
		case '/':
			return RIGHT
		case '\\':
			return LEFT
		}
	case DOWN:
		switch nextChar {
		case '.', '|':
			return DOWN
		case '-':
			return LEFTRIGHT
		case '/':
			return LEFT
		case '\\':
			return RIGHT
		}
	}
	panic("shouldn't be here")
}