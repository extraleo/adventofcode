package main

import (
	"extraleo/algorithm/utils"
	"fmt"
	"strconv"
)

func day() {
	sum := 0
	data, _ := utils.SpiltInputByLine(utils.INPUT)
	for _, item := range data {
		k, _ := strconv.Atoi(item)
		sum = sum + (k/3 - 2)
	}
	fmt.Println(sum)
}

func nightFunc(item int, sum int) (int, int) {
	if item/3-2 <= 0 {
		return item, sum
	}
	next := item/3 - 2
	sum = sum + next
	return nightFunc(next, sum)
}

func night() {
	sum := 0
	data, _ := utils.SpiltInputByLine(utils.INPUT)
	for _, item := range data {
		k, _ := strconv.Atoi(item)
		_, kSum := nightFunc(k, 0)
		sum = sum + kSum
	}
	fmt.Println(sum)
}

func main() {
	night()
}
