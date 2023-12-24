package main

import (
	"adventofcode/utils"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Lens struct{
	name string
	num uint8
}

type Box []Lens

func index(b Box, name string) int{
	for idx, l:=range b{
		if l.name == name{
			return idx
		}
	}
	return -1
}


func main() {
	// day()
	night()
}

func night() {
	var res = 0
	parts := parse()
	boxes:=make([]Box, 256)
	for _, part:=range parts{
		if (part[len(part)-1] == '-'){
			name := part[:len(part)-1]
			h := hashing(name)
			box := boxes[h]
			if i := index(box, name); i != -1 {
				boxes[h] = append(box[:i], box[i+1:]...)
			}
		}else if name, number, found := strings.Cut(part, "="); found {
			h := hashing(name)
			box := boxes[h]
			if i := index(box, name); i != -1 {
				boxes[h] = append(box[:i], Lens{name, uint8(utils.Atoi(number))})
				boxes[h] = append(boxes[h], box[i+1:]...)
			} else {
				boxes[h] = append(box, Lens{name, uint8(utils.Atoi(number))})
			}
		}
	}

	for bIdx, box:=range boxes{
		for i,l:=range box{
			res += (1 + bIdx) * (i + 1) * int(l.num)
		}
	}
	fmt.Println("night", res)
}

func day() {
	day := 0
	s := parse()
	for _, str := range s {
		day += hashing(str)
	}
	fmt.Println("day", day)
}

func parse() []string {
	input = strings.TrimSpace(input)
	steps := strings.Split(input, ",")
	return steps
}

func hashing(str string) int {
	hash := 0
	for _, c := range str {
		hash += int(c)
		hash *= 17
		hash %= 256
	}

	return hash
}


