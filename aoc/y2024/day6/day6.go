package main

import (
	. "adventofcode/utils/models"
	_ "embed"
	"fmt"
)

//go:embed sample.txt
var input string

var path map[Position]Position

var offsetStream = []Position{
	{X:-1,Y:0},
	{X:0,Y:1},
	{X:1,Y:0},
	{X:0,Y:-1},
}

func main(){
	day()
	night()
}

func night(){
	grid := ParseGrid(input)
	night := 0
	for su:=range(path){
		newGrid := copyAndSetGrid(grid, su)
		if ok, _ := trackGuard(newGrid);  !ok {
			night++
		}
	}
	fmt.Println("night: ", night)
}


func copyAndSetGrid(grid map[Position]rune,obstructions Position) map[Position]rune{
	copy := grid
	copy[obstructions] = '#'
	return copy
}


func day(){
	grid := ParseGrid(input)
	_, path = trackGuard(grid)
	fmt.Println("day: ",len(path))
}

func trackGuard(grid map[Position]rune) (bool, map[Position]Position){
	guard := getGuardPosition(grid)
	turnRight := 0
	offset := getOffset(turnRight, offsetStream)
	visited := make(map[Position]Position)
	visited[guard]=offset
	for ;; {
		nextGuard := nextGuard(guard, offset)
		if v, ok := grid[nextGuard]; ok{
			if v == '#'{
				turnRight++
				offset = getOffset(turnRight, offsetStream)
			}else{
				if existDirect, vOk:=visited[nextGuard]; !vOk {
					visited[nextGuard] = offset
				}else if existDirect == offset {
					return false, visited
				}
				guard = nextGuard
			}
		}else{
			return true, visited
		}
	}
}

func nextGuard(guard, offset Position) Position{
	guard.X = guard.X + offset.X
	guard.Y = guard.Y + offset.Y
	return guard
}

func getOffset(turnRight int, offsetStream []Position) Position{
	 p := offsetStream[turnRight%4]
	 return p
}



func getGuardPosition(grid map[Position]rune) Position{
	for k,v := range(grid){
		if v == '^'{
			return k	
		}
	}
	return Position{}
}