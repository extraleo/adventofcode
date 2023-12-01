package main

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func day(){
	re := regexp.MustCompile("[1-9]")
	result:=0
	input := utils.SpiltInputByLine("input.txt")
	for _,item := range input{
		numbers:=re.FindAllString(item, -1)
		firstNum,_:=strconv.Atoi(numbers[0])
		lastNum,_:=strconv.Atoi(numbers[len(numbers)-1])
		result=result + firstNum*10 + lastNum
	} 
	fmt.Printf("day: %d", result)
}

func night(){
	digits := make(map[string]int, 18)
	digits["1"]=1
	digits["one"]=1
	digits["2"]=2
	digits["two"]=2
	digits["3"]=3
	digits["three"]=3
	digits["4"]=4
	digits["four"]=4
	digits["5"]=5
	digits["five"]=5
	digits["6"]=6
	digits["six"]=6
	digits["7"]=7
	digits["seven"]=7
	digits["8"]=8
	digits["eight"]=8
	digits["9"]=9
	digits["nine"]=9

	re := regexp.MustCompile("[1-9]|one|two|three|four|five|six|seven|eight|nine")
	result:=0
	input := utils.SpiltInputByLine("input.txt")
	for _,item := range input{
		item=strings.Replace(item,"twone","2one",-1)
		item=strings.Replace(item,"oneight","1eight",-1)
		item=strings.Replace(item,"threeight","3eight",-1)
		item=strings.Replace(item,"fiveight","5eight",-1)
		item=strings.Replace(item,"sevenine","7nine",-1)
		item=strings.Replace(item,"eightwo","8two",-1)
		item=strings.Replace(item,"eighthree","8three",-1)
		item=strings.Replace(item,"nineight","9eight",-1)
		numbers:=re.FindAllString(item, -1)
		firstNum:=digits[numbers[0]]
		lastNum:=digits[numbers[len(numbers)-1]]
		result=result + firstNum*10 + lastNum
	} 
	fmt.Printf("night: %d", result)
}


func main(){
	item:="eighthree"	
	item=strings.Replace(item,"oneight","1eight",-1)

		item=strings.Replace(item,"threeight","3eight",-1)
		item=strings.Replace(item,"fiveight","5eight",-1)
		item=strings.Replace(item,"sevenine","7nine",-1)
		item=strings.Replace(item,"eightwo","8two",-1)
		item=strings.Replace(item,"eighthree","8three",-1)
		item=strings.Replace(item,"nineight","9eight",-1)
	re := regexp.MustCompile("([1-9]|one|two|three|four|five|six|seven|eight|nine)")
	fmt.Println(re.FindAllString(item,-1))
	night()
}