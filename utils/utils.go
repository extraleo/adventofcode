package utils

import (
	"io/ioutil"
	"log"
	"strings"
)

func SpiltInputByLine(inputPath string) ([]string, error) {
	return SpilInput(inputPath, "\n")
}

func SpilInput(inputPath string, sep string)([]string, error){
	data, err:=ioutil.ReadFile(inputPath)	
	if err!= nil{
		log.Fatal(err)
		return nil, err
	}
	return strings.Split(string(data),sep), err
}