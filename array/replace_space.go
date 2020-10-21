package main

import (
	"fmt"
	"strings"
)

// input: "Have a nice day"
// ountput: "Have%20a%20nice%20day"
// using join seems a little stupid, lets read Join code
func join(){
	inputStr:= strings.Split("Have a nice day"," ")
	output:=make([]string, 0, cap(inputStr)*2)
	for index, item:=range inputStr{
		if(index == len(inputStr)-1){
			output = append(output, item)
			break
		}
		output=append(output, item,"%20")
	}
	fmt.Println(strings.Join(inputStr,"%20"))
}


func main(){
	var test []int
	test = append(test, 9,4)
	fmt.Print(len(test))
}