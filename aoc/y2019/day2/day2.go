package main

import (
	"adventofcode/utils"
	"fmt"
)

func day(noun, verb int) int {
	data := utils.SpiltInputGetInt("input.txt", ",")
	data[1] = noun
	data[2] = verb
	source := extendAndCopy(data, len(data))
	for index := 0; index < len(source); {
		switch source[index] {
		case 1:
			source[source[index+3]] = source[source[index+1]] + source[source[index+2]]
			index = index + 4
		case 2:
			source[source[index+3]] = source[source[index+1]] * source[source[index+2]]
			index = index + 4
		case 99:
			return source[0]
		default:
			index++
		}
	}
	return -1
}

func extendAndCopy(source []int, len int) []int {
	target := make([]int, len+1)
	copy(target, source)
	return target
}

// seems only Brute-force search
func night() (int, int) {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb < 99; verb++ {
			if day(noun, verb) == 19690720 {
				return noun, verb
			}
		}
	}
	return 0, 0
}

func main() {
	noun, verb := night()
	fmt.Print(100*noun + verb)
}
