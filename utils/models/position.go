package models

import "adventofcode/utils"

type Position struct{
	X, Y int
}

func (pos *Position) Manhattan(p1 Position) int{
	return utils.Abs(pos.X - p1.X) + utils.Abs(pos.Y - p1.Y)
}

