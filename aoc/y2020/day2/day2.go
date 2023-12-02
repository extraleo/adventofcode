package main

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
)

func day() int {
	data := utils.SpiltInputByLine(utils.INPUT)
	count := 0
	for _, item := range data {
		aPSW := breakItem(item)
		repeatTime := strings.Split(aPSW[0], "-")
		min, _:=strconv.Atoi(repeatTime[0])
		 max,_:=strconv.Atoi(repeatTime[1])
		if validPWDNight(min,max, aPSW[1][0], aPSW[2]) {
			count++
		}
	}
	return count
}
func validPWD(minRepeatTime int, maxRepeatTime int, char byte, password string) bool {
	repeatCount := 0
	for i := 0; i < len(password); i++ {
		if char == password[i] {
			repeatCount++
		}
	}
	return repeatCount <= maxRepeatTime && repeatCount >= minRepeatTime
}
func validPWDNight(right int, left int, char byte, password string) bool{
		repeatCount := 0
		if password[right-1] ==  char{
			repeatCount ++
		}
		if password[left-1] == char{
			repeatCount++
		}
		return repeatCount == 1
	}


func breakItem(item string) []string {
	data := strings.Split(item, " ")
	return data
}

func main() {
	fmt.Print("Day: ", day())
}
