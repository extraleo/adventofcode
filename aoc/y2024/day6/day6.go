package main

import (
	. "adventofcode/utils/models"
	_ "embed"
	"fmt"
)

//go:embed sample.txt
var input string

func main(){
	day()
	night()
}

func day(){
	grid := ParseGrid(input)
	guard := getGuardPosition(grid)
	fmt.Println(guard)
}

func night(){

}

func getGuardPosition(grid map[Position]rune) Position{
	for k,v := range(grid){
		if v == '^'{
			return k	
		}
	}
	return Position{}
}