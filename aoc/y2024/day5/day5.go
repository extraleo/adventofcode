package main

import (
	_ "embed"
	"fmt"
)

//go:embed sample.txt
var input string

func day() {
	_day := 0
	fmt.Println("day:", _day)
}

func night(){
	_night :=0
	fmt.Println("night:", _night)
}

func main(){
	day()
	night()
}