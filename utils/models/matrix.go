package models

import (
	"strings"
)

type Matrix [][]rune

func ParseAsMatrixFromInput(str string) Matrix{
	str = strings.TrimSpace(str)
	lines := strings.Split(str, "\n")
	matrix := make(Matrix, len(lines))
	for idx, l := range lines {
		matrix[idx] = []rune(l)
	}
	return matrix
}

func (m Matrix) IsValidPosition(pos Position) bool{
	if pos.Y < 0 || pos.X < 0 {
		return false
	}
	if  pos.Y < m.LenY() && pos.X < m.LenX() {
		return true
	}
	return false
}

func (m Matrix) LenX() int{
	return len(m[0])
}

func (m Matrix) LenY() int{
	return len(m)
}

func (m Matrix) MaxX() int{
	if m.LenX() == 0 {
		return 0
	}
	return m.LenX() -1
}

func (m Matrix) MaxY() int{
	if m.LenY() == 0{
		return 0
	}
	return m.LenY() -1
}


func (m Matrix) IntValue(pos Position) int{
	return int(m[pos.Y][pos.X] - '0')
}