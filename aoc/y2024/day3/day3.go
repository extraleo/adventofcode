package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

func day() {
	sum := 0
	re := regexp.MustCompile(`mul\([0-9]{1,3}\,[0-9]{1,3}\)`)
	s := re.FindAllString(input, -1)
	reNum := regexp.MustCompile((`[0-9]{1,3}`))
	for _, i := range s {
		nums := reNum.FindAllString(i, -1)
		f, _ := strconv.Atoi(nums[0])
		s, _ := strconv.Atoi(nums[1])
		sum += f * s
	}
	fmt.Println(sum)
}

func night() {
	sum := 0
	re := regexp.MustCompile(`(do\(\)|don't\(\)|mul\([0-9]{1,3}\,[0-9]{1,3}\))`)
	s := re.FindAllString(input, -1)
	do := 1

	reNum := regexp.MustCompile((`[0-9]{1,3}`))
	reDo := regexp.MustCompile((`do\(\)`))
	reDont := regexp.MustCompile((`don't\(\)`))
	for _, i := range s {
		isDo := reDo.FindString(i)
		isDont := reDont.FindString(i)
		nums := reNum.FindAllString(i, -1)
		if (isDo != ""){
			do = 1
		}
		if (isDont != ""){
			do = 0
		}
		if (len(nums)!=0){
			f, _ := strconv.Atoi(nums[0])
			s, _ := strconv.Atoi(nums[1])
			sum += f * s * do
		}
	}
	fmt.Println(sum)
}

func main() {
	night()
}
