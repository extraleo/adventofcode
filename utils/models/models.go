package models

import "adventofcode/utils"


type ListNode struct{
	Val string
	Next *ListNode
}

type Point struct{
	X, Y int
}

type Grid map[Point]int

func (p *Point) Manhattan(p1 Point) int{
	return utils.Abs(p.X - p1.X) + utils.Abs(p.Y - p1.Y)
}


