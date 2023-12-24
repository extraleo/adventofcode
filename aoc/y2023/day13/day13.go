package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type pattern [][]rune

func main() {
	// day()
	night()
}

func night() {
	parts := parse()
	result := 0
	for _, part := range parts {
		idx, isHorizontal := findMirror(mapToCount(part))
		result += calculate(idx, isHorizontal)
	}
	fmt.Println("night", result)
}

func day() {
	parts := parse()
	result := 0
	for _, part := range parts {
		idx, isHorizontal := findMirror(mapToCount(part))
		result += calculate(idx, isHorizontal)
	}
	fmt.Println("day", result)
}

func getPattern(part string) pattern {
	lines := strings.Split(part, "\n")

	var pattern pattern
	for _, l := range lines {
		pattern = append(pattern, []rune(l))
	}
	return pattern
}

func parse() []pattern {
	parts := strings.Split(input, "\n\n")

	var patterns []pattern
	for _, b := range parts {
		patterns = append(patterns, getPattern(b))
	}
	return patterns
}

func findMirror(rowCnt, colCnt []int) (int, bool) {
	if r := findReflection(rowCnt); r >= 0 {
		return r + 1, true
	}
	return findReflection(colCnt) + 1, false
}

func findReflection(cnts []int) int {
	for i := 1; i < len(cnts); i++ {
		if isReflection(cnts, i-1, i) {
			return i - 1
		}
	}
	return -1
}

func isReflection(cnts []int, l, r int) bool {
	smudges:=0
	for l >= 0 && r < len(cnts){
		ok, smugde :=smudge(cnts[l], cnts[r])
		if (ok){
			smudges += smugde
			l--
			r++
		}else{
			break
		}
		
	}
	// if part1, remove smudges == 1 condition
	return (l < 0 || r == len(cnts)) && smudges == 1
}

func smudge(l, r int) (bool, int) {
	res := l ^ r
	if res == 0 {
		return true,0
	}
	if (res & (res - 1)) == 0 {
		return true,1
	}
	return false,0
}

func mapToCount(pat pattern) ([]int, []int) {
	rowCnt := make([]int, len(pat))
	colCnt := make([]int, len(pat[0]))

	for i, r := range pat {
		for j, v := range r {
			if v == '#' {
				rowCnt[i] |= 1 << j
				colCnt[j] |= 1 << i
			}
		}
	}
	/*
		Counts Result for example input
		RowCnt- [205 180 259 259 180 204 181] ColCnt- [77 12 115 33 82 82 33 115 12]
		RowCnt- [305 289 460 223 223 460 289] ColCnt- [91 24 60 60 25 67 60 60 103]
	*/
	return rowCnt, colCnt
}

func calculate(idx int, isHorizontal bool) int {
	if isHorizontal {
		return 100 * idx
	}
	return idx
}
