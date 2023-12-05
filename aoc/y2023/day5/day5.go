package main

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strings"
)

type innerMap map[int][]int


func night(){
    parts,_:=utils.SpiltInput("input.txt", "\n\n")
    re := regexp.MustCompile("[0-9]+")
    seeds :=re.FindAllString(parts[0], -1)
    reallySeeds:=reallySeed(seeds)
    almanac:=buildMappingList(parts)
    min:=0
    for i:=0; i<len(reallySeeds);{
        for seed:=reallySeeds[i];seed< reallySeeds[i]+reallySeeds[i+1]; seed++{
            
            location:=findLocation(seed, almanac)
            if i==0 || location < min {
                min=location
            }
        }
        
        i=i+2
        fmt.Print("iii",i)
    }
    fmt.Println("night", min)
}

func reallySeed(originSeeds []string) []int {
    count:=0
    result:=make([]int,len(originSeeds))
    for i := 0; i < len(originSeeds); i++ {
        result[i]=utils.Atoi(originSeeds[i])
        if(i%2==1){
            count+=utils.Atoi(originSeeds[i])
        }
    }
    fmt.Println("will count:",count)
    return result
}

func day(){
    parts,_:=utils.SpiltInput("input.txt", "\n\n")
    for _,part:=range parts{
        fmt.Println(part)
    }
    re := regexp.MustCompile("[0-9]+")
    seeds :=re.FindAllString(parts[0], -1)
    almanac:=buildMappingList(parts)
    min:=0
    for i, seed:= range seeds{
        fmt.Println("seed:", utils.Atoi(seed))
        location:=findLocation(utils.Atoi(seed), almanac)
        if i==0 || location < min {
         min=location
     }

    }
    fmt.Println("day", min)
    
}

func findLocation(seed int, almanac map[string]innerMap) int{
    
    ss:=almanac["seed-to-soil"]
    soil:=findNext(seed, ss)
    

    sf:=almanac["soil-to-fertilizer"]
    fer:=findNext(soil,sf)
    fw:=almanac["fertilizer-to-water"]
    water:=findNext(fer, fw)
    wl:=almanac["water-to-light"]
    light:=findNext(water, wl)
    lt:=almanac["light-to-temperature"]
    temper:=findNext(light, lt)
    th:=almanac["temperature-to-humidity"]
    hum:=findNext(temper, th)
    hl:=almanac["humidity-to-location"]
    return findNext(hum, hl)
    

}
func findNext(key int, almanacInnerMap innerMap) int {
    for source,v:= range almanacInnerMap{
        if key >= source && key <= source+v[1]-1{
            result:= v[0] + (key - source)
            return result
        }
    }
    return key
}




func buildMappingList(parts []string) map[string]innerMap{
    result := make(map[string]innerMap)
    for i := 1; i < len(parts); i++ {
        part:=strings.Split(parts[i], "\n")
        mapKey:=strings.Split(part[0], " ")[0]
        innerMap:=make(innerMap)
        for k:=1;k< len(part); k++{
            t:=strings.Split(part[k], " ")
            innerMap[utils.Atoi(t[1])]=[]int{utils.Atoi(t[0]),utils.Atoi(t[2])}
        }
        result[mapKey]=innerMap
    }
    return result
}

func main(){
    night()
}