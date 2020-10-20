package utils

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const INPUT="input.txt"

func SpiltInputByLine(inputPath string) ([]string, error) {
	return SpilInput(inputPath, "\n")
}

func SpiltInputGetInt(inputPath string, sep string)([]int, error)  {
	data, err := SpilInput(inputPath, sep)
	if err != nil {
		return nil, err
	}
	result := make([]int, len(data))
  for index,item :=range data{
		result[index],_=strconv.Atoi(item)
	}
  return result, err

}

func SpilInput(inputPath string, sep string)([]string, error){
	data, err:=ioutil.ReadFile(inputPath)	
	if err!= nil{
		log.Fatal(err)
		return nil, err
	}
	return strings.Split(string(data),sep), err
}