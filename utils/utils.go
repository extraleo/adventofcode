package utils

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func SpiltInputByLine(inputPath string) []string {
	data, err := SpiltInput(inputPath, "\n")
	if err != nil {
		panic(err)
	}
	return data
}

func SpiltInputGetInt(inputPath string, sep string) []int {
	data, err := SpiltInput(inputPath, sep)
	if err != nil {
		panic(err)
	}
	result := make([]int, len(data))
	for index, item := range data {
		result[index], _ = strconv.Atoi(item)
	}
	return result

}

func SpiltInputGetIntList(input string, sepLine string) [][]int {
	lines := strings.Split(input, "\n")
	result := [][]int{}
	for _, line := range lines {
		lineToken := strings.Split(line, sepLine)
		nums := []int{}
		for _, i := range lineToken {
			nums = append(nums, Atoi(i))
		}
		result = append(result, nums)
	}
	return result

}

func SpiltInput(inputPath string, sep string) ([]string, error) {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return strings.Split(string(data), sep), err
}

func ConvertIntString2IntList(str string) []int {
	re := regexp.MustCompile("[0-9]+")

	s := re.FindAllString(str, -1)
	result := make([]int, len(s))
	for i, e := range s {
		result[i] = Atoi(e)
	}
	return result

}

func ConvertIntString2Int8List(str string) []uint8 {
	re := regexp.MustCompile("[0-9]+")

	s := re.FindAllString(str, -1)
	result := make([]uint8, len(s))
	for i, e := range s {
		result[i] = uint8(Atoi(e))
	}
	return result

}

func ConvertIntString2String(str string) []string {
	re := regexp.MustCompile("[0-9]+")
	return re.FindAllString(str, -1)
}

func Atoi(s string) int {
	item, _ := strconv.Atoi(s)
	return item
}

type MatrixChar [][]rune

func getPattern(part string) MatrixChar {
	lines := strings.Split(part, "\n")

	var pattern MatrixChar
	for _, l := range lines {
		pattern = append(pattern, []rune(l))
	}
	return pattern
}

func ParseAsMatrixChars(input string) []MatrixChar {
	parts := strings.Split(input, "\n\n")

	var patterns []MatrixChar
	for _, b := range parts {
		patterns = append(patterns, getPattern(b))
	}
	return patterns
}
