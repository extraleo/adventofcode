package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	day()
	night()
}

const dot = 46

func night() {
	result := 0
	gearMap:=make(map[string]int, 14000)
	lines := utils.SpiltInputByLine("input.txt")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		for j := 0; j < len(line); j++ {
			if isNum(line[j]) {
				startIndex, endIndex := j, j
				for endIndex < len(line) && isNum(line[endIndex]) {
					endIndex++
				}
				if startIndex > 0 {
					startIndex--
				}
				if endIndex == len(line){
					endIndex--
				}
				partNum, gear :=isPartNum(startIndex, endIndex ,i, lines)
				if partNum && gear != "" {
					part:=convertNum(startIndex, endIndex, line)
					if(gearMap[gear] == 0){
						gearMap[gear]=part
					}else{
						firstGear := gearMap[gear]
						result = result + firstGear * part
					}
				}
				j = endIndex
			}
		}
	}
	fmt.Println("night:", result)
}

func day() {
	result := 0
	lines := utils.SpiltInputByLine("input.txt")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		for j := 0; j < len(line); j++ {
			if isNum(line[j]) {
				startIndex, endIndex := j, j
				for endIndex < len(line) && isNum(line[endIndex]) {
					endIndex++
				}
				if startIndex > 0 {
					startIndex--
				}
				if endIndex == len(line){
					endIndex--
				}
				partNum, _ :=isPartNum(startIndex, endIndex ,i, lines)
				if partNum {
					result= result + convertNum(startIndex, endIndex, line)
				}
				j = endIndex
			}
		}
	}
	fmt.Println("day:", result)
}

func isNum(s byte) bool{
	return s >= 48 && s <= 57
}

func convertNum(start, end int, line string) int {
	if !isNum(line[start]) {
		start++
	}
	if isNum(line[end]) {
		end++
	}
	return utils.Atoi(line[start:end])
}


func isPartNum(start, end, row int, lines []string ) (bool, string) {
	var currentLine, upLine, downLine string
	currentLine = lines[row]
	if (row == 0){
		upLine = ""
		downLine = lines[row+1]
	}else if(row == len(lines)-1){
		upLine = lines[row-1]
		downLine = ""
	}else{
		upLine = lines[row-1]
		downLine = lines[row+1]
	}
	// check current
	if (!isNum(currentLine[start]) && currentLine[start] != dot) {
		if (currentLine[start] == '*'){
			return true, fmt.Sprintf("%d-%d", row, start)
		}
		return true, ""
	}
	if (!isNum(currentLine[end]) && currentLine[end] != dot){
		if (currentLine[end] == '*'){
			return true, fmt.Sprintf("%d-%d", row, end)
		}
		return true, ""
	}
	// check up line
	if upLine != "" {
		for i := start; i <= end; i++ {
			if !isNum(upLine[i]) && upLine[i] != dot {
				if (upLine[i]=='*'){
					return true, fmt.Sprintf("%d-%d", row-1, i)
				}
				return true,""
			}
		}
	}

	// check down
	if downLine != "" {
		for i := start; i <= end; i++ {
			if !isNum(downLine[i]) && downLine[i] != dot {
				if (downLine[i] == '*'){
					return true, fmt.Sprintf("%d-%d", row+1, i)
				}
				return true, ""
			}
		}
	}

	return false, ""
}