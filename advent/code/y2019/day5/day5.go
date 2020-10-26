package main

import "extraleo/algorithm/utils"

func day(noun, verb int) int {
	data, _ := utils.SpiltInputGetInt(utils.INPUT, ",")
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

func night(){

}

func main()  {
	
}