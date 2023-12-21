package main

import _ "embed"

//go:embed input-test.txt
var input string

type pattern [][]rune

func main() {
	day()
	night()
}

func night() {
	panic("unimplemented")
}

func day() {
	panic("unimplemented")
}

func parse() []pattern {
	res := make([]pattern, 0)
	return res
}

func findMirror(pat pattern) (int, bool) {
	return -1, false
}

func findVertical(pat pattern) int {
	return -1
}

func findHorizontal(pat pattern) int {
	return -1
}

func calculate(idx int, isHorizontal bool) int {
	if isHorizontal {
		return 100 * (idx + 1)
	}
	return idx + 1
}
