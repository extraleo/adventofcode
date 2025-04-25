package main

import (
	"adventofcode/leetcode/utils"
	"fmt"
)

func lengthOfLongestSubstring(s string) int {

  start := 0 
	max := 0
	vMap:=make(map[byte]int, len(s))
	for i:=0; i< len(s); i++{
		if v, ok := vMap[s[i]]; ok {
			if v < start {
				max = utils.MaxInt(max, i-start+1)
				vMap[s[i]] = i
				continue
			}
			max = utils.MaxInt(max, i-start)
			start = v + 1
			vMap[s[i]] = i
			continue
		}
		vMap[s[i]] = i
		max = utils.MaxInt(max, i-start+1)
	}
	return max
}

func main(){
	fmt.Println(lengthOfLongestSubstring("aabaab!bb"))
}