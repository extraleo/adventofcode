package main

import (
	"adventofcode/utils"
	_ "embed"
	"strings"
)

//go:embed input-test.txt
var input string


func day(){
parts := strings.Split(input, "\n\n")
instruction:=parts[0]
lines := strings.Split(parts[1], "\n")
treeMap:=make(map[string]utils.Tree)
for _,line:=range lines{
	tokens:=strings.Split(line, " = (")
	LR:=strings.Split()
}
}

func night(){

}

func main(){

}