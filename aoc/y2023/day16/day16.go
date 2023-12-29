package main

import (
	"adventofcode/utils/models"
	_ "embed"
)

//go:embed input-test.txt
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
	direction int
	position models.Position
}


func main() {
	day()
	night()
}

func night() {
	panic("unimplemented")
}

func day() {
	matrix := models.ParseAsMatrixFromInput(input)
	start
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