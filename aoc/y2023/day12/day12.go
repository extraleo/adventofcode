package main

import (
	"adventofcode/utils"
	_ "embed"
	"fmt"
	"strings"
)

// thanks https://www.reddit.com/r/adventofcode/comments/18ghux0/comment/kd0npmi/?utm_source=share&utm_medium=web2x&context=3
func main() {
	// day()
	night()
}

//go:embed input.txt
var input string

type springInfo struct {
	spring  string
	numbers string
}

var cache = make(map[springInfo]int)

func setCache(spring string, numbers []uint8, value int) int {
	cache[springInfo{spring, string(numbers)}] = value
	return value
}

func count(spring string, numbers []uint8) int {
	if len(spring) == 0 && len(numbers) == 0 {
		return 1
	}

	if len(spring) == 0 {
		return 0
	}

	if value, ok := cache[springInfo{spring, string(numbers)}]; ok {
		return value
	}

	// cut dot
	if spring[0] == '.' {
		res := count(spring[1:], numbers)
		return setCache(spring, numbers, res)
	}

	if spring[0] == '?' {
		res := count(spring[1:], numbers) + count("#"+spring[1:], numbers)
		return setCache(spring, numbers, res)
	}

	if spring[0] == '#' {
		if len(numbers) == 0 {
			res := 0
			return setCache(spring, numbers, res)
		}

		number := numbers[0]
		// find first dot
		// only two branch:
		// index < number => return 0
		// index >= number ==> count(spring[number:], numbers[1:])  like: #???.######..#####. 1,6,5
		indexDot := strings.Index(spring, ".")
		if indexDot == -1 {
			indexDot = len(spring)
		}
		if indexDot < int(number) {
			res := 0
			return setCache(spring, numbers, res)
		}
		// like: #???.######..#####. 1,6,5
		// remain = ???.######..#####

		remaining := spring[number:]
		if len(remaining) == 0 {
			res := count(remaining, numbers[1:])
			return setCache(spring, numbers, res)
		}

		if remaining[0] == '#' {
			// fail
			res := 0
			return setCache(spring, numbers, res)
		}

		res := count(remaining[1:], numbers[1:])
		return setCache(spring, numbers, res)
	}
	panic("cannot happen")
}

func anotherCount(spring string, numbers []uint8) int {
	if len(spring) == 0 && len(numbers) == 0 {
		return 1
	}

	if len(spring) == 0 {
		return 0
	}

	if value, ok := cache[springInfo{spring, string(numbers)}]; ok {
		return value
	}

	// cut dot
	if spring[0] == '.' {
		res := count(spring[1:], numbers)
		return setCache(spring, numbers, res)
	}

	if spring[0] == '?' {
		res := count(spring[1:], numbers) + count("#"+spring[1:], numbers)
		return setCache(spring, numbers, res)
	}

	if spring[0] == '#' {
		// out of bound
		if len(numbers) == 0 {
			res := 0
			return setCache(spring, numbers, res)
		}

		// check length >= sum number
		sum := 0
		for _, n := range numbers {
			sum += int(n)
		}
		if len(spring) < sum {
			res := 0
			return setCache(spring, numbers, res)
		}

		// check first group
		// all chars in firs group are ? or #
		number := numbers[0]
		for i := 0; i < int(number); i++ {
			if spring[i] == '.' {
				res := 0
				return setCache(spring, numbers, res)
			}
		}

		res := count(spring[number+1:], numbers[1:])
		return setCache(spring, numbers, res)
	}
	panic("ERROR!!!!!")
}

func day() {

	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	result := 0
	for _, line := range lines {
		fields := strings.Split(line, " ")
		springs := fields[0]
		numbers := utils.ConvertIntString2Int8List(fields[1])
		result += anotherCount(springs, numbers)
	}

	fmt.Println("day:", result)

}

func unfoldSpring(spring string) string {
	tmp := "?" + spring
	for i := 0; i < 4; i++ {
		spring = spring + tmp
	}
	return spring
}

func unfoldNumbers(numbers []uint8) []uint8 {
	var res []uint8
	for i := 0; i < 5; i++ {
		res = append(res, numbers...)
	}
	return res
}

func night() {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	result := 0
	for _, line := range lines {
		fields := strings.Split(line, " ")
		springs := fields[0]
		numbers := utils.ConvertIntString2Int8List(fields[1])
		springs = unfoldSpring(springs)
		numbers = unfoldNumbers(numbers)
		result += anotherCount(springs, numbers)
	}

	fmt.Println("night:", result)
}
