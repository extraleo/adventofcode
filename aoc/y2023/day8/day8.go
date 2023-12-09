package main

import (
	"adventofcode/utils"
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

type point struct {
	left  string
	right string
}

func day() {
	result := 0
	parts := strings.Split(input, "\n\n")
	instruction := parts[0]
	lines := strings.Split(parts[1], "\n")
	network := make(map[string]point)
	re := regexp.MustCompile("[A-Z]+")
	for _, line := range lines {
		tokens := re.FindAllString(line, -1)
		network[tokens[0]] = point{tokens[1], tokens[2]}
	}

	current := "AAA"
	for current != "ZZZ" {
		index := result % len(instruction)
		step := instruction[index]
		p := network[current]
		if step == 'L' {
			current = p.left
		} else {
			current = p.right
		}
		result++
	}

	fmt.Println("day:", result)

}

func night() {
	results := []int{}
	parts := strings.Split(input, "\n\n")
	instruction := parts[0]
	lines := strings.Split(parts[1], "\n")
	network := make(map[string]point)
	starts := []string{}
	re := regexp.MustCompile("[A-Z]+")
	for _, line := range lines {
		tokens := re.FindAllString(line, -1)
		network[tokens[0]] = point{tokens[1], tokens[2]}
		if tokens[0][2] == 'A' {
			starts = append(starts, tokens[0])
		}
	}
	for _, start := range starts {
		result := 0
		current := start
		for current[2] != 'Z' {
			index := result % len(instruction)
			step := instruction[index]
			p := network[current]
			if step == 'L' {
				current = p.left
			} else {
				current = p.right
			}
			result++
		}
		results = append(results, result)
	}
	night := utils.LCM(results[0], results[1], results...)
	fmt.Println("night:", night)

}

func main() {
	night()
}
