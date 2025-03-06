package models

import (
	"adventofcode/utils"
	"strings"
)

type Position struct{
	X, Y int
}


func (pos *Position) Manhattan(p1 Position) int{
	return utils.Abs(pos.X - p1.X) + utils.Abs(pos.Y - p1.Y)
}




func ParseGrid(str string) map[Position]rune{
	lines:=strings.Split(str, "\n")
	result := make(map[Position]rune)
	for i,l:=range(lines){
		for j,r:=range(l){
				p := Position{X: i,Y: j}
				result[p] = r
			}
		}
	return result
}
