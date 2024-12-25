package main

import (
	"adventofcode/utils"
	_ "embed"
	"fmt"
	"math"
)

//go:embed input.txt
var input string

func day() {
	reports := utils.SpiltInputGetIntList(input, " ")
	sum := 0
	for _, report := range reports {
		if safe, _ := isSafeDay(report); safe {
			sum++
		}
	}
	fmt.Println("day: ", sum)
}

func isSafeDay(report []int) (bool, int) {
	if report[len(report)-1]-report[0] == 0 {
		return false, 0
	}
	if report[len(report)-1]-report[0] > 0 {
		for i := range report {
			if i != len(report)-1 && (report[i+1]-report[i] <= 0 || report[i+1]-report[i] > 3) {
				return false, i+1
			}
		}
		return true, 0
	}
	for i := range report {
		if i != len(report)-1 && (report[i+1]-report[i] >= 0 || report[i+1]-report[i] < -3) {
			return false, i+1
		}
	}
	return true, 0

}

func isSafeNight(report []int) bool {
	if safe, i := isSafeDay(report); !safe {
		ret := make([]int, 0)
		ret = append(ret, report[:i]...)
		ret = append(ret, report[i+1:]...)
		safeNight, _ := isSafeDay(ret)
		if safeNight {
			// fmt.Println("before: ", report, i)
			// fmt.Println("after: ", ret)
		}else{
			fmt.Println("before: ", report, i)
			fmt.Println("after: ", ret)
		}
		return safeNight

	}
	return true

}

func compareLevels(report []int) int {
	initialDiff := report[1] - report[0]
	safe := 1
	for i := 0; i < len(report)-1; i++ {
		difference := report[i+1] - report[i]
		if difference*initialDiff <= 0 || int(math.Abs(float64(difference))) > 3 {
			safe = 0
			break
		}
	}
	return safe
}

func compareLevelsWithDampener(report []int) int {
	safe := compareLevels(report)
	if safe == 0 {
		for i := range report {
			reportCopy := arrayCopy(report)
			newReport := append(reportCopy[:i], reportCopy[i+1:]...)
			oneRemovedSafe := compareLevels(newReport)
			if oneRemovedSafe == 1 {
				return 1
			}
		}
	}
	return safe
}



func arrayCopy(orig []int) []int {
	newNums := make([]int, len(orig))
	copy(newNums, orig[:])
	return newNums
}

func night() {
	reports := utils.SpiltInputGetIntList(input, " ")
	sum := 0
	for _, report := range reports {
		sum +=compareLevelsWithDampener(report)
	}
	fmt.Print("night: ", sum)
}

func main() {
	// day()
	night()
}
