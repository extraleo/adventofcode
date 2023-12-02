package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func day(){
	redCubes:=12
	greenCubes:=13
	blueCubes:=14
	games := utils.SpiltInputByLine("input.txt")
	gameIds:= make([]string, len(games))

	for _, game := range games{
		s := strings.Split(game, ":")
		gameId := strings.Split(s[0], " ")[1]
		gameRounds := strings.Split(s[1], ";")
		canPlay:=true
		for _,round :=range gameRounds{
			roundDetail :=strings.Split(round, ",")
			for _,cube :=range roundDetail{
				cubeDetail := strings.Split(cube, " ")
				if ((cubeDetail[2] == "red" && utils.Atoi(cubeDetail[1]) > redCubes) ||
				 		(cubeDetail[2] == "green" && utils.Atoi(cubeDetail[1]) > greenCubes) ||
						(cubeDetail[2] == "blue" && utils.Atoi(cubeDetail[1]) > blueCubes)){
					canPlay = false
				}
			}
		}
		if (canPlay){
			gameIds = append(gameIds, gameId)
		}
	}
	result := 0
	for _,id:=range gameIds{
		result = result + utils.Atoi(id)
	}
	fmt.Println("day result:", result)
}

func night(){
	result:=0
	games := utils.SpiltInputByLine("input.txt")
	for _, game := range games{
		s := strings.Split(game, ":")
		gameRounds := strings.Split(s[1], ";")
		redMax:=0
		greenMax:=0
		blueMax:=0
		for _,round :=range gameRounds{
			roundDetail :=strings.Split(round, ",")
			for _,cube :=range roundDetail{
				cubeDetail := strings.Split(cube, " ")
				if (cubeDetail[2] == "red" && utils.Atoi(cubeDetail[1]) > redMax){
					redMax = utils.Atoi(cubeDetail[1])
				} 
				if (cubeDetail[2] == "green" && utils.Atoi(cubeDetail[1]) > greenMax){
					greenMax = utils.Atoi(cubeDetail[1])
				}	
				if(cubeDetail[2] == "blue" && utils.Atoi(cubeDetail[1]) > blueMax){
					blueMax = utils.Atoi(cubeDetail[1])
				}
			}
		}
		result = result + redMax*greenMax*blueMax
	}
	fmt.Println("night result:", result)
}



func main(){
	day()
	night()
}