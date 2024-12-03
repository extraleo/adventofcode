package main

import (
	"adventofcode/utils"
	"fmt"
	"sort"
	"strings"
)

func helper() ([]int, []int) {
	lines := utils.SpiltInputByLine("input.txt")
	var first, second []int
	for _, line := range lines {
		l := strings.Fields(line)
		first = append(first, utils.Atoi(l[0]))
		second = append(second, utils.Atoi(l[1]))
	}
	return first, second
}

func day() {
	f, s := helper()
	sort.Ints(f)
	sort.Ints(s)
	sum := 0
	for i, _ := range f {
		if f[i] >= s[i] {

			sum += f[i] - s[i]
		} else {
			sum += s[i] - f[i]

		}
	}
	fmt.Println(sum)

}

func night() {
	f, s := helper()
	times := make(map[int]int)
	for _, i := range s {
		times[i] += 1
	}
	sum := 0
	for _, j := range f {
		if v, ok := times[j]; ok {
			sum += v * j
		}
	}
	print(sum)
}

func main() {
	night()
}
