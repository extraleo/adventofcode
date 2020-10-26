package main

import (
	"extraleo/algorithm/utils"
)

type point struct {
	x int
	y int
}

type vector struct {
	start point
	end   point
}

func isIntersect(v1, v2 vector) bool {
	return false
}

func getIntersectionPoint(v1, v2 vector) point {
	return point{1, 2}
}

func getDistance(p point) int {
	return utils.Abs(p.x) + utils.Abs(p.y)
}

func getVectors(wires []string) []vector {
	
	return []vector{}
}

func day() int {
	data := utils.SpiltInputByLine(utils.INPUT)
	wires1, _ := utils.SpilInput(data[0], ",")
	wires2, _ := utils.SpilInput(data[1], ",")
	vectors1 := getVectors(wires1)
	vectors2 := getVectors(wires2)
	var manhatten []int
	for i := 0; i < len(vectors1); i++ {
		for k := 0; k < len(vectors2); k++ {
			if isIntersect(vectors1[i], vectors2[k]) {
				manhatten = append(manhatten, getDistance(getIntersectionPoint(vectors1[i], vectors2[k])))
			}
		}
	}
	return utils.Min(manhatten...)
}

func night() {

}

func main() {

}
