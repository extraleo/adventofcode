package main

import "fmt"

func main() {
	// 避免负数, 直接散列用下标
	nums := [4]int{-1, 1, 2, 3}
	s := 0
	for n := range nums {
		s = 1<<n + s
	}
	resultTotal := make([][]int, 0, len(nums))
	for sub := s; ; {
		sub = (sub - 1) & s
		subList := make([]int, 0, len(nums))
		for i := range nums {
			if (sub>>i)&1 == 1 {
				subList = append(subList, nums[i])
			}
		}
		resultTotal = append(resultTotal, subList)
		if sub == s {
			break
		}
	}
	fmt.Println("xxx:", resultTotal)
}
