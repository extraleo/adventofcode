package main

import (
	"fmt"
	"strconv"
)

// my input is 271973 and 785961
func day() (int, int) {
	count := 0
	nightCount := 0
	for i := 271973; i < 785961; i++ {
		if fact1(i) && fact2(i) {
			count++
			if fact3(i) {
				nightCount++
			}
		}
	}
	return count, nightCount
}

func fact1(i int) bool {
	str := strconv.Itoa(i)
	for k := 0; k < len(str)-1; k++ {
		if str[k+1] == str[k] {
			return true
		}
	}
	return false
}

func fact2(i int) bool {
	str := strconv.Itoa(i)
	for k := 0; k < len(str)-1; k++ {
		if str[k+1] < str[k] {
			return false
		}
	}
	return true
}

func fact3(i int) bool {
	s := strconv.Itoa(i)
	characters := []rune(s)
	seen := make(map[rune]int)
	for _, v := range characters {
		seen[v] = seen[v] + 1
	}
	for _, v := range seen {
		if v == 2 {
			return true
		}
	}
	return false
}

func main() {
	d, n := day()
	fmt.Printf("Day %d, night %d", d, n)
}
