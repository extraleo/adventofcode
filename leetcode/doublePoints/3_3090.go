package main

import (
	"adventofcode/leetcode/utils"
	// "fmt"
)

func lengthOfLongestSubstring(s string) int {
	return maxLengthSubstrings(s, 1)
}

func maximumLengthSubstring(s string) int {
	return maxLengthSubstrings(s, 2)
}

// m == repeat count
func maxLengthSubstrings(s string, m int) int {
  start, max := 0, 0
	vMap := map[byte]int{}
	for i := 0; i< len(s); i++{
		// if exist
		if v, ok := vMap[s[i]]; ok {
			// plus 1
			vMap[s[i]] = v + 1
			if vMap[s[i]] >= m {
				vMap[s[start]] = vMap[s[start]] -1
				max = utils.MaxInt(max, i-start)
				start++
				continue
			}
		}
		vMap[s[i]] = 1
		max = utils.MaxInt(max, i-start+1)
	}
	return max
}

// deprecated
// 3 solution
func firstEdition(s string) int {
	start := 0
	max := 0
	vMap := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
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




// func main() {
// 	fmt.Println(lengthOfLongestSubstring("pwwkew"))
// }
