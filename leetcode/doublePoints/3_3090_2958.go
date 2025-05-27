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
  left, ans := 0, 0
	vMap := map[byte]int{}
	for i:=range len(s) {
		v := s[i]
		vMap[v] = vMap[v] + 1
		for vMap[v] > m{
			vMap[s[left]] = vMap[s[left]] - 1
			left++
		}
		ans = max(ans, i-left+1)
	}
	return ans
}

// deprecated
// 3 solution
// 这个版本不适用于 k =2 的情况, 主要是是出在如何找到第二个重复字符的位置
// 在最初的版本上我用了一个额外的数组来记录重复字母的位置
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
