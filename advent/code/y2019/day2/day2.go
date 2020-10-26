package main

import (
	"extraleo/algorithm/utils"
	"fmt"
)

func day(noun, verb int) int {
	data := utils.SpiltInputGetInt(utils.INPUT, ",")
	data[1] = noun
	data[2] = verb
	source := extendAndCopy(data, len(data))
	for index := 0; index < len(source); {
		switch source[index] {
		case 1:
			// if source[index+3] >= len(source) {
			// 	source = extendAndCopy(source, source[index+3])
			// }
			// if max(source[index+1], source[index+2], 0) >= len(source) {
			// 	source = extendAndCopy(source,  max(source[index+1], source[index+2], 0))
			// }
			source[source[index+3]] = source[source[index+1]] + source[source[index+2]]
			index = index + 4
		case 2:
			// if source[index+3] >= len(source) {
			// 	source = extendAndCopy(source, source[index+3])
			// }
			// if max(source[index+1], source[index+2], 0) >= len(source) {
			// 	source = extendAndCopy(source,max(source[index+1], source[index+2], 0))
			// }
			source[source[index+3]] = source[source[index+1]] * source[source[index+2]]
			index = index + 4
		case 99:
			// fmt.Println("The result: ", source[0])
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
