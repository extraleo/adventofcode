package main

import (
	"adventofcode/utils"
	_ "embed"
	"strconv"
	"strings"
)

//go:embed sample.txt
var input string

func main(){
	day()
	night()
}

func day(){
	day := 0
	calis := parseInput(input)
	for cal, values := range(calis) {
			if (operatorCheck(cal, values)){
				day += cal
			}
	}
}

func night(){

}

func parseInput(input string) map[int][]int{
	lines := strings.Split(input, "\n")
	re := make(map[int][]int)
	for _, line := range(lines) {
		l := strings.Split(line, ":")
		calibration, _ :=strconv.Atoi(l[0])
		re[calibration] = utils.ConvertIntString2IntList(l[1])
	}
	return re
}

func operatorCheck(calibration int, values []int) bool {
	if (!minMaxCheck(calibration, values)){
		return false
	}
	
}


func minMaxCheck(calibration int, values []int) bool{
	plusResult :=0 
	multiplyResult := 1
	for _, i := range(values){
		plusResult += i
		multiplyResult *= i
	}
	if plusResult > calibration || multiplyResult < calibration {
		return false
	}
	return true
}



