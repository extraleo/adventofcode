package utils

import (
	"io/ioutil"
	"log"
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

func SpiltInput(inputPath string, sep string) ([]string, error) {
	data, err := ioutil.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return strings.Split(string(data), sep), err
}


func ConvertIntString2IntList(str string) []int{
	re:=regexp.MustCompile("[0-9]+")
	
	s := re.FindAllString(str, -1)
	result:=make([]int, len(s))
	for i,e:=range s{
		result[i]=Atoi(e)
	}
	return result

}


func ConvertIntString2String(str string) []string{
	re:=regexp.MustCompile("[0-9]+")
	return re.FindAllString(str, -1)
}



func Atoi(s string) int {
	item, _ := strconv.Atoi(s)
	return item
}
