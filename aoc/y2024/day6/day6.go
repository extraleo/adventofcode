package main

import (
	. "adventofcode/utils/models"
	_ "embed"
	"fmt"
)

//go:embed sample.txt
var input string

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



func day(){
	grid := ParseGrid(input)
	repeat, pathLen := trackGuard(grid)
	fmt.Println("day: ",repeat, len(pathLen))
}

func trackGuard(grid map[Position]rune) (int, map[Position]Position){
	guard := getGuardPosition(grid)
	turnRight := 0
	offset := getOffset(turnRight, offsetStream)
	visited := make(map[Position]Position)
	visited[guard]=offset
	repeat := 0
	for ;; {
		nextGuard := nextGuard(guard, offset)
		if v, ok := grid[nextGuard]; ok{
			if v == '#'{
				turnRight++
				offset = getOffset(turnRight, offsetStream)
			}else{
				if offsetVisited, vOk:=visited[nextGuard]; !vOk {
					visited[nextGuard] = offset
				}else{
					if offsetVisited != offset{
						repeat++
						fmt.Println(nextGuard, offset)
					}
				}
				guard = nextGuard
			}
		}else{
			return repeat, visited
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

func night(){
	fmt.Println("night: ")
}

func getGuardPosition(grid map[Position]rune) Position{
	for k,v := range(grid){
		if v == '^'{
			return k	
		}
	}
	return Position{}
}