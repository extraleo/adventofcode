package utils

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const INPUT = "input.txt"

func SpiltInputByLine(inputPath string) []string {
	data, err := SpilInput(inputPath, "\n")
	if err != nil {
		panic(err)
	}
	return data
}

func SpiltInputGetInt(inputPath string, sep string) []int {
	data, err := SpilInput(inputPath, sep)
	if err != nil {
		panic(err)
	}
	result := make([]int, len(data))
	for index, item := range data {
		result[index], _ = strconv.Atoi(item)
	}
	return result

}

func SpilInput(inputPath string, sep string) ([]string, error) {
	data, err := ioutil.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return strings.Split(string(data), sep), err
}

func Max(item ...int) int {
	max := 0
	for _, i := range item {
		if i > max {
			max = i
		}
	}
	return max
}

func Min(item ...int) int {
	min := 0
	for _, i := range item {
		if i < min {
			min = i
		}
	}
	return min
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Atoi(s string) int {
	item, _ := strconv.Atoi(s)
	return item
}
