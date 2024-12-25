package main

import (
	"adventofcode/utils/models"
	"fmt"
)

//go:embed sample.txt
var input string

func day()  {
	_day := 0
	matrix := models.ParseAsMatrixFromInput(input)
	for i :=0; i< matrix.MaxX(); i++{
		for j:=0; j < matrix.MaxY(); j++{
			if matrix[i][j] == 'X' || matrix[i][j] == 'S' {
				_day += startSearch(i,j)
			}
		}
	}
	fmt.Println("day:", _day)

}

func startSearch(matrix models.Matrix,i,j int ) int {
	
	return 0
}



func main(){
	day()
}