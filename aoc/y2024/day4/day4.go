package main

import (
	"adventofcode/utils/models"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func day() {
	_day := 0
	matrix := models.ParseAsMatrixFromInput(input)
	for i := 0; i <= matrix.MaxY(); i++ {
		for j := 0; j <= matrix.MaxX(); j++ {
			if matrix[i][j] == 'X' {
				_day += startSearchX(matrix, i, j)
			}
		}
	}
	fmt.Println("day:", _day)

}

func night() {
	_night := 0
	matrix := models.ParseAsMatrixFromInput(input)
	for i := 0; i <= matrix.MaxY(); i++ {
		for j := 0; j <= matrix.MaxX(); j++ {
			if matrix[i][j] == 'A' {
				_night += mSearchD(matrix, i, j)
			}
		}
	}
	fmt.Println("night:", _night)

}

func startSearchX(matrix models.Matrix, i, j int) int {
	h := xSearchH(matrix, i, j)
	v := xSearchV(matrix, i, j)
	d := xSearchD(matrix, i, j)
	return h + v + d
}
func xSearchH(matrix models.Matrix, i, j int) int {

	h := 0
	if j+3 <= matrix.MaxX() {
		if matrix[i][j+1] == 'M' && matrix[i][j+2] == 'A' && matrix[i][j+3] == 'S' {
			fmt.Println("h+", i, j)
			h++
		}
	}
	if j-3 >= 0 {
		if matrix[i][j-1] == 'M' && matrix[i][j-2] == 'A' && matrix[i][j-3] == 'S' {
			fmt.Println("h-", i, j)
			h++
		}
	}
	return h
}

func xSearchV(matrix models.Matrix, i, j int) int {
	v := 0
	if i+3 <= matrix.MaxY() {
		if matrix[i+1][j] == 'M' && matrix[i+2][j] == 'A' && matrix[i+3][j] == 'S' {
			fmt.Println("v+", i, j)
			v++
		}
	}
	if i-3 >= 0 {
		if matrix[i-1][j] == 'M' && matrix[i-2][j] == 'A' && matrix[i-3][j] == 'S' {
			fmt.Println("v-", i, j)
			v++
		}
	}
	return v
}

func xSearchD(matrix models.Matrix, i, j int) int {
	d := 0
	if i+3 <= matrix.MaxY() && j+3 <= matrix.MaxX() {
		if matrix[i+1][j+1] == 'M' && matrix[i+2][j+2] == 'A' && matrix[i+3][j+3] == 'S' {
			fmt.Println("d++", i, j)
			d++
		}
	}
	if i-3 >= 0 && j-3 >= 0 {
		if matrix[i-1][j-1] == 'M' && matrix[i-2][j-2] == 'A' && matrix[i-3][j-3] == 'S' {
			fmt.Println("d--", i, j)

			d++
		}
	}
	if i-3 >= 0 && j+3 <= matrix.MaxX() {
		if matrix[i-1][j+1] == 'M' && matrix[i-2][j+2] == 'A' && matrix[i-3][j+3] == 'S' {
			fmt.Println("d-+", i, j)

			d++
		}
	}
	if i+3 <= matrix.MaxY() && j-3 >= 0 {
		if matrix[i+1][j-1] == 'M' && matrix[i+2][j-2] == 'A' && matrix[i+3][j-3] == 'S' {
			fmt.Println("d+-", i, j)

			d++
		}
	}
	return d
}


func mSearchD(matrix models.Matrix, i, j int) int{
	if i + 1 <= matrix.MaxY() && i -1 >=0 && j+1 <= matrix.MaxX() && j-1>=0 {
		if (matrix[i+1][j+1] == matrix[i+1][j-1] && matrix[i-1][j+1] == matrix[i-1][j-1]){
			if (matrix[i+1][j+1] == 'S' && matrix[i-1][j-1] == 'M') ||(matrix[i+1][j+1] == 'M' && matrix[i-1][j-1] == 'S'){
				return 1
			}
		}
		if (matrix[i+1][j-1] == matrix[i-1][j-1] && matrix[i-1][j+1] == matrix[i+1][j+1]){
			if (matrix[i+1][j-1] == 'S' && matrix[i-1][j+1] == 'M') ||(matrix[i+1][j-1] == 'M' && matrix[i-1][j+1] == 'S'){
				return 1
			}
		}
	}

	return 0
}

func main() {
	night()
}
