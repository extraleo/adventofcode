package main

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strings"

	"github.com/bits-and-blooms/bitset"
)

func day() {
	cards := utils.SpiltInputByLine("input.txt")
	re := regexp.MustCompile("[0-9]+")
	result := 0
	for _, card := range cards {
		winning, having := &bitset.BitSet{}, &bitset.BitSet{}
		cardDetail := strings.Split(card, "|")
		bitSetAddStringList(re.FindAllString(strings.Split(cardDetail[0], ":")[1], -1), winning)
		bitSetAddStringList(re.FindAllString(cardDetail[1], -1), having)
		if winning.Intersection(having).Count() > 0 {
			result = result + 1<<(winning.Intersection(having).Count()-1)
		}
	}
	fmt.Println("day:", result)
}

func night() {
	cards := utils.SpiltInputByLine("input.txt")
	re := regexp.MustCompile("[0-9]+")
	result := make([]int, len(cards))
	for i := 0; i < len(cards); i++ {
		result[i] = 1
	}
	for index, card := range cards {
		winning, having := &bitset.BitSet{}, &bitset.BitSet{}
		cardDetail := strings.Split(card, "|")
		bitSetAddStringList(re.FindAllString(strings.Split(cardDetail[0], ":")[1], -1), winning)
		bitSetAddStringList(re.FindAllString(cardDetail[1], -1), having)
		winCards := int(winning.Intersection(having).Count())
		for k := 0; k < result[index]; k++ {
			for i := index + 1; i <= index+winCards; i++ {
				result[i] = result[i] + 1
			}
		}
	}
	total := 0
	for i := 0; i < len(cards); i++ {
		total = total + result[i]
	}
	fmt.Println("night:", total)
}

func bitSetAddStringList(origin []string, b *bitset.BitSet) {
	for _, s := range origin {
		b.Set(uint(utils.Atoi(s)))
	}
}

func main() {
	night()
}
