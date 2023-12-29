package main

import (
	"adventofcode/utils/models"
	"adventofcode/utils/set"
	_ "embed"
	"fmt"

	"golang.org/x/text/currency"
)

//go:embed input-test.txt
var input string

func main(){
	day()
	night()
}

func night() {
	panic("unimplemented")
}

func day() {
	day:=0
  grid := models.ParseAsMatrixFromInput(input)
	candicate :=[]models.Position
	visited := set.NewSet[models.Position]()
	results := []models.Position{models.Position{X:0, Y:0}}
	current := results[len(results)-1] 

	for current.X != grid.MaxX() && current.Y != grid.LenY(){
		
	}


	for _,r:=range results{
		day += grid.IntValue(r)
	}
	fmt.Println("day", day)
}
