package main

import (
	"adventofcode/utils"
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input-test.txt
var input string

type Hand struct {
	card string
	bit  int
}

func compareHand(h1, h2 Hand) int {
	if h1.card == h2.card {
		return 0
	}
	h2Type, h1Type := h2.typeOfCard(), h1.typeOfCard()
	if h2Type-h1Type == 0 {
		for i := range h2.card {
			if index(h2.card[i]) == index(h1.card[i]) {
				continue
			}
			return index(h1.card[i]) - index(h2.card[i])
		}
	}
	return h1Type - h2Type
}

const (
	HIGHT = iota
	PAIR
	TWO_PAIRS
	THREE_KIND
	FULL_HOUSE
	FOUR_KIND
	FIVE_KIND
)

func (h Hand) typeOfCard() int {
	tmp := make([]int, 15)
	for i := 0; i < len(h.card); i++ {
		tmp[index(h.card[i])]++
	}

	jCount:= tmp[index('J')]
	if joker {
		tmp[index('J')] = 0
	}

	slices.SortFunc(tmp, func(a, b int) int { return b - a })

	if (joker){
		tmp[0]+=jCount
	}

	if tmp[0] == 5 {
		return FIVE_KIND
	}

	if tmp[0] == 4 {
		return FOUR_KIND
	}

	if tmp[0] == 3 && tmp[1] == 2 {
		return FULL_HOUSE
	}
	if tmp[0] == 3 {
		return THREE_KIND
	}
	if tmp[0] == 2 && tmp[1] == 2 {
		return TWO_PAIRS
	}
	if tmp[0] == 2 {
		return PAIR
	}
	return HIGHT

}

var jValue int
var joker bool

func index(c byte) int {
	if joker{
		jValue = 1
	}else{
		jValue =11
	}

	if c >= '2' && c <= '9' {
		return int(c - '0')
	}
	switch c {
	case 'T':
		return 10
	case 'J':
		return jValue
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	}
	panic("unknown word")
}

func day() {
	joker =false
	s:=utils.SpiltInputByLine("input.txt")
	hands:=make([]Hand, 0,len(s))
	for _,l:=range s{
		token := strings.Split(l, " ")
		hands = append(
			hands,
			Hand{
				card: token[0],
				bit: utils.Atoi(token[1]),
			},
		)
	}
	fmt.Println("hands:", hands, len(hands))
	slices.SortFunc(hands, func(h1, h2 Hand) int {return compareHand(h1, h2 )})
	result:=0
	for i,h:=range hands{
		result = result + (i+1)*h.bit
	}
	fmt.Println("day:", result)
}

func night() {
	joker =true
	s:=utils.SpiltInputByLine("input.txt")
	hands:=make([]Hand, 0,len(s))
	for _,l:=range s{
		token := strings.Split(l, " ")
		hands = append(
			hands,
			Hand{
				card: token[0],
				bit: utils.Atoi(token[1]),
			},
		)
	}
	slices.SortFunc(hands, func(h1, h2 Hand) int {return compareHand(h1, h2)})
	result:=0
	for i,h:=range hands{
		result = result + (i+1)*h.bit
	}
	fmt.Println("night:", result)
}

func main() {
	night()
}
