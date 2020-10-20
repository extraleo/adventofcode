package main

import (
	"extraleo/algorithm/utils"
	"fmt"
)

func day(){
	data, _ := utils.SpiltInputGetInt(utils.INPUT, ",")
	data[1] = 12
	data[2] = 2
	source := extendAndCopy(data, len(data))
	for index:=0; index < len(source); {
		switch source[index] {
		case 1:
			if source[index+3] >= len(source) {
				source = extendAndCopy(source, source[index+3])
			}
			if max(source[index+1], source[index+2], 0) >= len(source) {
				source = extendAndCopy(source,  max(source[index+1], source[index+2], 0))
			}
			source[source[index+3]] = source[source[index+1]] + source[source[index+2]]
			index=index+4
		case 2:
			if source[index+3] >= len(source) {
				source = extendAndCopy(source, source[index+3])
			}
			if max(source[index+1], source[index+2], 0) >= len(source) {
				source = extendAndCopy(source,max(source[index+1], source[index+2], 0))
			}
			source[source[index+3]] = source[source[index+1]] * source[source[index+2]]
			index=index+4
		case 99:
			fmt.Println("The result: ", source[0])
			return
		default:
			index++
		}
	}
}

func extendAndCopy(source []int, len int) []int {
	target := make([]int, len+1)
	copy(target, source)
	return target
}
func max(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		} else {
			return c
		}
	} else {
		if b > c {
			return b
		} else {
			return c
		}
	}
}

func night() {

}

func main() {
	day()
}
